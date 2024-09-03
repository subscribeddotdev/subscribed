package wait_for

import (
	"os"
	"sync"
	"time"

	_ "github.com/lib/pq"
	"github.com/rabbitmq/amqp091-go"
	"github.com/subscribeddotdev/subscribed/server/internal/common/logs"
	"github.com/subscribeddotdev/subscribed/server/internal/common/postgres"
)

type WaitFor struct {
	wg     *sync.WaitGroup
	logger *logs.Logger
}

func NewWaitFor(logger *logs.Logger) *WaitFor {
	return &WaitFor{
		logger: logger,
		wg:     &sync.WaitGroup{},
	}
}

func (w *WaitFor) Do(handler func() error, svcName string, timeout time.Duration) {
	w.wg.Add(1)

	w.logger.Info("⏱️ WaitFor started", "service", svcName)

	go func() {
		until := time.NewTimer(timeout)

		for {
			select {
			case <-until.C:
				w.logger.Error("⛔  Timeout reached", "service", svcName)
				w.wg.Done()

				// Exit with an error because the checks must be all or nothing.
				os.Exit(1)
			default:
				err := handler()
				if err != nil {
					w.logger.Debug("⚠️ Check failed, trying again...in 250ms", "service", svcName, "error", err)
					time.Sleep(time.Millisecond * 250)
					continue
				}

				if err == nil {
					w.logger.Info("✅  Ready", "service", svcName)
					w.wg.Done()
					return
				}
			}
		}
	}()
}

func (w *WaitFor) Wait() {
	w.wg.Wait()
}

// Run Wait for other containers to start responding before running the service
func Run() {
	w := NewWaitFor(logs.New())
	w.Do(func() error {
		db, err := postgres.Connect(os.Getenv("DATABASE_URL"))
		if err != nil {
			return err
		}

		return db.Ping()
	}, "postgres", time.Second*30)

	w.Do(func() error {
		_, err := amqp091.Dial(os.Getenv("AMQP_URL"))
		if err != nil {
			return err
		}

		return nil
	}, "rabbitmq", time.Second*30)

	w.Wait()
}
