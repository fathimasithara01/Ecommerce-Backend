package services

import (
	"errors"
	"fmt"

	repository "github.com/fathimasithara01/ecommerce/src/repository/user"
	"github.com/fathimasithara01/ecommerce/utils/models"
)

func ProcessPayment(userID uint, req models.PaymentRequest) error {
	order, err := repository.GetOrderById(req.OrderID, userID)
	if err != nil || order.UserID != userID {
		return errors.New("invalid order")
	}

	// Step 3: Check if the order is already marked as paid
	isPaid, err := repository.IsOrderPaid(order.ID)
	if err != nil {
		return fmt.Errorf("failed to check payment status: %w", err)
	}
	if isPaid {
		return errors.New("order already paid")
	}

	var status string
	var paymentID string

	switch req.PaymentMode {
	case "COD":
		status = "pending"
	case "RAZORPAY":
		paymentID = "razorpay_txn_123"
		status = "success"
	case "STRIPE":
		paymentID = "stripe_txn_456"
		status = "success"
	default:
		return errors.New("invalid payment method")
	}

	payment := models.Payment{
		OrderID:       order.ID,
		PaymentMethod: req.PaymentMode,
		Status:        status,
		// Amount:      order.TotalAmount,
		TransactionID: paymentID,
	}

	if err := repository.SavePayment(&payment); err != nil {
		return err
	}

	return repository.UpdateOrderStatus(order.ID, status)
}
