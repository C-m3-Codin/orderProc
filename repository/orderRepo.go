package repository

import (
	"c-m3-codin/ordProc/models"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepo struct {
	DB *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepo {

	or := OrderRepo{
		DB: db,
	}

	return or
}

func (orderRep OrderRepo) GetOrder(orderId string) (order models.Order, err error) {
	order.ID, err = uuid.Parse(orderId)
	if err != nil {
		return
	}
	// fmt.Println(order)
	err = orderRep.DB.First(&order).Error

	return
}

func (orderRep OrderRepo) CreateOrder(order models.Order) (err error) {

	result := orderRep.DB.Create(&order)
	err = result.Error
	return
}

func (OrderRepo OrderRepo) UpdateOrder(order models.Order) (err error) {
	err = OrderRepo.DB.Model(&models.Order{}).Where("id = ?", order.ID).Updates(models.Order{Status: 2, OrderCompleted: order.OrderCompleted}).Error
	return
}

func (OrderRepo OrderRepo) GetUnproccessedOrders() (orders []models.Order, err error) {
	err = OrderRepo.DB.Where("status < ? ", 2).Find(&orders).Error
	return
}

func (o OrderRepo) GetProccessedCount() (count int64, err error) {
	fmt.Println("c here GetProccessedCount")
	if err := o.DB.Model(&models.Order{}).Where("status = ?", 1).Count(&count).Error; err != nil {
		fmt.Println("Error getting processed orders count: %v", err)
	}
	fmt.Println("count here GetProccessedCount", count)
	return
}

func (o OrderRepo) GetCompletedCount() (count int64, err error) {
	if err := o.DB.Model(&models.Order{}).Where("status = ?", 2).Count(&count).Error; err != nil {
		fmt.Println("Error getting processed orders count: %v", err)
	}
	return
}

func (o OrderRepo) GetTotalCount() (count int64, err error) {
	// assuming total order doesnt include pending orders
	fmt.Println("c here GetTotalCount")
	if err := o.DB.Model(&models.Order{}).Count(&count).Error; err != nil {
		log.Fatalf("Error counting total orders: %v", err)
	}
	fmt.Println("count GetTotalCount ", count)
	return
}

// type respo struct {
// 	TotalProcessingTime time.Duration `gorm:"column:total_processing_time"`
// 	OrderCount          int64         `gorm:"column:order_count"`
// }

func (o OrderRepo) GetAverageProcessingTimeCount() (totalProcessingTime time.Duration, err error) {
	// var r respo
	var totalProcessingTimeString string
	if err := o.DB.Model(&models.Order{}).
		Where("status = ? AND order_completed IS NOT NULL AND order_received IS NOT NULL", 2).
		Select("SUM(order_completed - order_received)").
		Scan(&totalProcessingTimeString).Error; err != nil {
		log.Fatalf("Error getting average processing time: %v", err)
	}
	fmt.Println("processing time total", totalProcessingTimeString)

	totalProcessingTime = convertToDuration(totalProcessingTimeString)
	if err != nil {
		log.Fatalf("Error parsing total processing time: %v", err)
	}

	// Calculate the average processing time if there are any orders

	// if r.OrderCount > 0 {
	// 	metrics.AverageProcessingTime = r.TotalProcessingTime / time.Duration(r.OrderCount)
	// 	fmt.Printf("Average processing time: %v\n", metrics.AverageProcessingTime)
	// } else {
	// 	fmt.Println("No processed orders found.")
	// }
	fmt.Println("processing time total in time", totalProcessingTime)

	return
}

func convertToDuration(durationString string) (duration time.Duration) {
	parts := strings.Split(durationString, ":")
	if len(parts) != 3 {
		log.Fatalf("Unexpected format of duration string: %v", durationString)
	}

	// Parse hours, minutes, and the second + microseconds part
	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatalf("Error parsing hours: %v", err)
	}

	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("Error parsing minutes: %v", err)
	}

	// Parse seconds and microseconds
	secondsAndMicroseconds := parts[2]
	secondsParts := strings.Split(secondsAndMicroseconds, ".")
	if len(secondsParts) != 2 {
		log.Fatalf("Unexpected format for seconds: %v", secondsAndMicroseconds)
	}

	// Parse seconds and microseconds
	seconds, err := strconv.Atoi(secondsParts[0])
	if err != nil {
		log.Fatalf("Error parsing seconds: %v", err)
	}

	microseconds, err := strconv.Atoi(secondsParts[1])
	if err != nil {
		log.Fatalf("Error parsing microseconds: %v", err)
	}

	// Convert to time.Duration
	duration = time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute +
		time.Duration(seconds)*time.Second + time.Duration(microseconds)*time.Microsecond
	return
}
