package main

// import "time"

const testVersion = 4

type cell struct {
  value int
  ch chan int // out chan
  quit chan string
}

// type inputCell cell;

func (cell cell) SetValue(newValue int) {
  cell.value = newValue

  // ch := make(chan int)
  cell.ch <- newValue
  <- cell.quit
  // go func () {
  //   cell.ch <- newValue

  //   for {
  //     select {
  //       // case "received" == <- cell.quit:
  //       case <- cell.quit:
  //         return
  //       case <- time.After(time.Millisecond):
  //         continue
  //     }
  //   }
  // }()

  // return ch
}

func (cell cell) Value() int {
  return cell.value
}

func (cell cell) getChannel() chan int {
  return cell.ch
}


// Reactor

type reactor int

func New() reactor {
  var dd reactor
  return dd
}

func (react *reactor) CreateInput(value int) cell {
  return cell{value, make(chan int), make(chan string)}
}

// (react *reactor) func CreateCompute1(cell Cell, clb func(int) int) ComputeCell {
func (react *reactor) CreateCompute1(ce cell, clb func(int) int) ComputeCell {
  // ch := cell.getChannel()
  // ch := ce.ch
  compCell := computeCell{&cell{value: clb(ce.Value()), ch: make(chan int)}}

  go func () {
    for {
      compCell.cell.value = clb(<- ce.ch)
      ce.quit <- "received"
    }
  }()

  return compCell
}

// type computeCell cell
type computeCell struct {
  cell *cell
}

func (compCell computeCell) Value() int {
  return compCell.cell.value;
}