# twilio-otp-backend
A simple, clean, and scalable OTP sending system using **Go worker pool + job queue + Twilio API**.   The app demonstrates how to send OTP messages asynchronously using buffered channels and background workers.

# Twilio OTP Queue Worker (Golang)

A simple, clean, and scalable OTP sending system using **Go worker pool + job queue + Twilio API**.  
The app demonstrates how to send OTP messages asynchronously using buffered channels and background workers.

This project is perfect for learning:
- Go concurrency (goroutines + channels)
- Worker pattern
- Message queue in Go
- Integrating Twilio SMS API
- Basic OTP job processing

---

## 🚀 Features

- Queue-based OTP job system  
- Background worker (goroutine) for processing jobs  
- Twilio SMS API integration  
- Non-blocking OTP enqueueing  
- dotenv support for configuration  
- Clean project structure  
- Realistic production-like job handling

---

## 📁 Project Structure

twilio-otp-app/
│── model/
│ └── task.go # Task struct (phone + otp)
│
│── queue/
│ └── queue.go # Job queue (buffered channel)
│
│── worker/
│ └── worker.go # Background job processor
│
│── main.go # App entry point
│── .env # Twilio config
│── go.mod
│── README.md
