package solution

import (
	"fmt"
	"sync"
	"time"

	"github.com/adesupraptolaia/bank-account-test/helper"
	"github.com/adesupraptolaia/bank-account-test/model"
)

type worker struct {
	wg        sync.WaitGroup
	mutex     sync.Mutex
	iteration int
}

func BeforeEodSolution1() {
	records := helper.ReadCSVFile("before_eod.csv")

	var w1, w2, w3 worker

	for i := 1; i <= 4; i++ {
		w1.wg.Add(1)
		go w1.threadOne(records, i)
	}

	for i := 1; i <= 4; i++ {
		w2.wg.Add(1)
		go w2.threadTwo(records, i)
	}

	for i := 1; i <= 8; i++ {
		w3.wg.Add(1)
		go w3.threadThree(records, i)
	}

	w1.wg.Wait()
	w2.wg.Wait()
	w3.wg.Wait()

	helper.WriteToCSV("after_eod.csv", records)
}

func (w *worker) threadOne(records []model.AfterEOD, threadNo int) {
	defer w.wg.Done()

	for {
		w.mutex.Lock()
		i := w.iteration
		w.iteration++
		w.mutex.Unlock()

		if i >= 200 {
			break
		}

		records[i].AverageBalance = (records[i].Balance + records[i].PreviousBalance) / 2
		records[i].ThreadNo1 = fmt.Sprintf("1-%d", threadNo)
		time.Sleep(1 * time.Microsecond)
	}
}

func (w *worker) threadTwo(records []model.AfterEOD, threadNo int) {
	defer w.wg.Done()

	for {
		w.mutex.Lock()
		i := w.iteration
		w.iteration++
		w.mutex.Unlock()

		if i >= 200 {
			break
		}

		if records[i].Balance > 150 {
			records[i].Balance += 25
			records[i].ThreadNo2b = fmt.Sprintf("2b-%d", threadNo)
		} else if records[i].Balance >= 100 {
			records[i].FreeTransfer = 5
			records[i].ThreadNo2a = fmt.Sprintf("2a-%d", threadNo)
		}
		time.Sleep(1 * time.Microsecond)
	}
}

func (w *worker) threadThree(records []model.AfterEOD, threadNo int) {
	defer w.wg.Done()

	for {
		w.mutex.Lock()
		i := w.iteration
		w.iteration++
		w.mutex.Unlock()

		if i >= 100 {
			break
		}

		records[i].Balance += 10
		records[i].ThreadNo3 = fmt.Sprintf("3-%d", threadNo)
		time.Sleep(1 * time.Microsecond)
	}
}
