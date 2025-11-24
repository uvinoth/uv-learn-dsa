package api

// Min Stack implemention
type MinStack struct {
	stack    []int
	minStack []int
}

func NewMinStack() *MinStack {
	return &MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (ms *MinStack) Push(val int) {
	ms.stack = append(ms.stack, val)

	if len(ms.minStack) == 0 || val <= ms.minStack[len(ms.minStack)-1] {
		ms.minStack = append(ms.minStack, val)
	}
}

func (ms *MinStack) Pop() {
	if len(ms.stack) == 0 {
		return
	}

	top := ms.stack[len(ms.stack)-1]
	ms.stack = ms.stack[:len(ms.stack)-1]

	if top == ms.minStack[len(ms.minStack)-1] {
		ms.minStack = ms.minStack[:len(ms.minStack)-1]
	}
}

func (ms *MinStack) Top() (int, bool) {
	if len(ms.stack) == 0 {
		return 0, false
	}
	return ms.stack[len(ms.stack)-1], true
}

func (ms *MinStack) GetMin() (int, bool) {
	if len(ms.minStack) == 0 {
		return 0, false
	}
	return ms.minStack[len(ms.minStack)-1], true
}

var ms = NewMinStack()

// func main() {
// 	router := gin.Default()

// 	// Push
// 	router.POST("/push", func(c *gin.Context) {
// 		var body struct {
// 			Value int `json:"value"`
// 		}

// 		if err := c.ShoulBindJSON(&body); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
// 		}

// 		ms.Push(body.Value)

// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "Pushed",
// 			"Value":   body.Value,
// 		})

// 	})

// 	// Pop
// 	router.DELETE("/pop", func(c *gin.Context) {
// 		ms.Pop()
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "Popped top element",
// 		})
// 	})

// 	// Get Top
// 	router.GET("top", func(c *gin.Context) {
// 		val, ok := ms.Top()
// 		if !ok {
// 			c.JSON(http.StatusNotFound, gin.H{
// 				"message": "Stack empty",
// 			})
// 		}
// 		c.JSON(http.StatusOK, gin.H{
// 			"value": val,
// 		})
// 	})

// 	// Get Min Element
// 	router.GET("/min", func(c *gin.Context) {
// 		val, ok := ms.GetMin()
// 		if !ok {
// 			c.JSON(http.StatusNotFound, gin.H{
// 				"message": "Stack empty",
// 			})
// 		}
// 		c.JSON(http.StatusOK, gin.H{
// 			"min": val,
// 		})
// 	})

// 	router.Run(":8080")

// }
