# Golang Internship 2025

## Guide

### Create the new project

- Create new repository from GitHub and clone the repository to your local machine
- Initialize Go module inside the folder: `go mod init github.com/yourusername/golang-internship-2025`
- Create project structure with proper package organization

### Setup Project Structure

```
golang-internship-2025/
├── .github/
│   └── workflows/
│       └── ci.yml
├── basic/
├── string/
├── array/
├── concurrency/
├── oop/
└── README.md
```

### Config GitHub Actions for CI/CD

Create `.github/workflows/ci.yml`:

```yaml
name: Go CI

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: "1.21"

      - name: Build
        run: go build -v ./...

      - name: Test
        run: go test -v -race -coverprofile=coverage.out ./...

      - name: Upload coverage to Codecov
        uses: codecov/codecov-action@v3
        with:
          token: ${{ secrets.CODECOV_TOKEN }}
          files: ./coverage.out
          fail_ci_if_error: true
```

### Config Codecov

- Create a Codecov account and link it to your GitHub account
- Create a new repository in Codecov and link it to your GitHub repository
- Go to Settings → Add Repository Secret in GitHub
- Add `CODECOV_TOKEN` with the token from Codecov

### Create the badges

Add to README.md:

```markdown
[![Go CI](https://github.com/yourusername/golang-internship-2025/actions/workflows/ci.yml/badge.svg)](https://github.com/yourusername/golang-internship-2025/actions/workflows/ci.yml)
[![codecov](https://codecov.io/github/yourusername/golang-internship-2025/branch/main/graph/badge.svg)](https://codecov.io/github/yourusername/golang-internship-2025)
[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/golang-internship-2025)](https://goreportcard.com/report/github.com/yourusername/golang-internship-2025)
```

### Test everything works

- Make changes to the code, commit and push to the repository
- Wait for the build to complete, refresh the repository to observe the changes
- Check Go Report Card for code quality metrics

## Testing Guidelines

- All functions must have corresponding unit tests with filename `*_test.go`
- Use table-driven tests for multiple test cases
- Aim for at least 80% code coverage
- Use `go test -v ./...` to run all tests
- Use `go test -cover ./...` to check coverage

Example test structure:

```go
func TestFunctionName(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected int
    }{
        {"case 1", 5, 10},
        {"case 2", 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := FunctionName(tt.input)
            if result != tt.expected {
                t.Errorf("got %v, want %v", result, tt.expected)
            }
        })
    }
}
```

## Exercises

### Basic

1. Viết chương trình tìm ước số chung lớn nhất (GCD) và bội số chung nhỏ nhất (LCM) của hai số tự nhiên
2. Viết chương trình tìm tổng các chữ số của một số nguyên
3. Viết chương trình nhận vào một số nguyên và trả về cách phân tích số đó ra tích của thừa số nguyên tố. Ví dụ nhập vào 600 thì cần phải trả về `2 * 2 * 2 * 3 * 5 * 5`
4. Số Fibonacci là dãy số bắt đầu từ 1, 1 và sau đó, số tiếp theo bằng 2 số trước cộng lại. Tức là `1 1 2 3 5 8 ...`. Nhập vào số nguyên n, in ra danh sách những số Fibonacci không lớn hơn n.
5. Viết hàm tính giai thừa `n!` của một số (implement cả recursive và iterative)
6. Viết chương trình nhận vào một số nguyên từ 1-12, trả về tên của tháng tương ứng
7. Số 23 khi viết nhị phân sẽ là 10111, khi viết ngược lại thành 11101 và nó thành số 29. Nhập số nguyên n và tìm số m bằng cách viết ngược thứ tự số binary như trên.
8. Viết chương trình dịch số thập phân thành số La Mã (1-3999)

### String

1. Kiểm tra chuỗi ký tự có đối xứng hay không (palindrome). Ví dụ: "abcdcba" là đối xứng
2. Tính tổng của các số nguyên trong chuỗi. Ví dụ: "abc 123 def 33 mn 3.221" sẽ in ra 380 với 380 = 123+33+3+221
3. Viết chương trình kiểm tra một chuỗi string có phải là sự lặp lại của một chuỗi con bất kỳ nào đó hay không. Ví dụ: "abcabcabc" là lặp lại của "abc"
4. Nhận vào một chuỗi ký tự chứa toàn các chữ cái (a-z, A-Z). Rút gọn chuỗi bằng cách ở những nơi chữ cái lặp lại, ta viết số lần lặp. Ví dụ: "abcccceeeeeefd" sẽ viết thành "abc4e6fd", "abbbbbbbbbbbbbc" viết là "ab13c"
5. Đảo ngược bài trên, ví dụ nhận vào "ab13c" trả về "abbbbbbbbbbbbbc"
6. Một Barcode EAN-13 có 13 con số được coi là hợp lệ nếu: tổng của các số ở vị trí lẻ + 3*(tổng các số ở vị trí chẵn) là một số chia hết cho 10. Ví dụ mã barcode 8938505974194 ta có (8+3+5+5+7+1+4) + 3*(9+8+0+9+4+9) = 150. 150 chia hết cho 10 nên mã 8938505974194 là hợp lệ. Viết chương trình kiểm tra tính hợp lệ của một barcode.
7. Viết chương trình đếm số lần xuất hiện của mỗi ký tự trong một chuỗi (sử dụng map)

### Slice & Array

1. Viết chương trình tìm tổng của các số chẵn và trừ đi tổng các số lẻ trong một slice số nguyên
2. Viết chương trình sử dụng phương pháp tìm kiếm nhị phân để tìm kiếm một số nguyên từ một slice số nguyên đã sắp xếp (tăng dần hoặc giảm dần), trả về vị trí nếu tìm thấy hoặc -1 nếu không tìm thấy
3. Viết chương trình đảo ngược các phần tử trong một slice số nguyên (yêu cầu: in-place, không sử dụng slice phụ)
4. Viết chương trình nhận vào 2 slice các số nguyên đã sắp xếp, tạo slice số nguyên từ 2 slice số nguyên đã cho và giữ nguyên thứ tự đã sắp xếp. (Yêu cầu: không dùng thêm thao tác sort)
5. Một slice số nguyên n phần tử có giá trị từ 1 đến n, bị bỏ đi một giá trị bất kỳ rồi sắp xếp ngẫu nhiên. Viết chương trình nhận vào slice số nguyên của quá trình trên và truy tìm số đã bị bỏ đi.
6. Viết chương trình tìm phần tử lớn thứ k trong một slice không sắp xếp (không dùng sort toàn bộ slice)
7. Viết chương trình xoay (rotate) slice sang trái hoặc phải k vị trí

### Maps & Structs

1. Viết chương trình đếm tần suất xuất hiện của các từ trong một đoạn văn bản
2. Tạo struct `Person` với các field: Name, Age, Email. Viết các method để:
   - Validate email hợp lệ
   - Kiểm tra tuổi có hợp lệ (0-150)
   - In thông tin người dùng theo format đẹp
3. Tạo struct `Student` embedding struct `Person`, thêm field GPA. Implement method tính xếp loại (Excellent: GPA >= 3.6, Good: GPA >= 3.2, Average: GPA >= 2.0, Poor: GPA < 2.0)
4. Viết chương trình quản lý danh bạ điện thoại với các chức năng:
   - Thêm contact mới
   - Tìm contact theo tên
   - Xóa contact
   - Liệt kê tất cả contacts sắp xếp theo tên
5. Implement cache đơn giản với map, có các method Get, Set, và Delete. Thêm tính năng TTL (time to live) cho mỗi entry.

### Interfaces & Polymorphism

1. Tạo interface `Shape` với method `Area()` và `Perimeter()`. Implement cho các struct: `Circle`, `Rectangle`, `Triangle`
2. Tạo interface `Sortable` và implement các method để sort một slice của custom type
3. Viết chương trình có interface `PaymentMethod` với method `Pay(amount float64) error`. Implement cho các struct: `CreditCard`, `Cash`, `MobilePayment`
4. Tạo interface `Storage` với methods `Save()`, `Load()`, `Delete()`. Implement cho `FileStorage` và `MemoryStorage`
5. Implement pattern Strategy: tạo interface `CompressionStrategy` và các implementation: `ZipCompression`, `RarCompression`, `GzipCompression`

### Error Handling

1. Viết chương trình đọc file và xử lý các error có thể xảy ra (file không tồn tại, không có quyền đọc, etc.)
2. Tạo custom error types cho một hệ thống banking với các errors: `InsufficientFundsError`, `InvalidAccountError`, `TransactionError`
3. Implement retry mechanism: viết hàm thực hiện một operation và tự động retry nếu gặp error (với backoff exponential)
4. Viết chương trình xử lý panic và recover để program không crash
5. Implement error wrapping: tạo chain of errors với context thông tin khi error đi qua các layers

### Concurrency (Goroutines & Channels)

1. Viết chương trình tính tổng các số từ 1 đến N sử dụng multiple goroutines
2. Implement worker pool pattern: tạo pool của workers để xử lý jobs từ một queue
3. Viết chương trình download multiple URLs concurrently và aggregate kết quả
4. Implement producer-consumer pattern sử dụng channels
5. Tạo rate limiter sử dụng channels để giới hạn số requests per second
6. Viết chương trình tìm số nguyên tố trong một range sử dụng concurrent goroutines
7. Implement timeout cho một operation sử dụng `select` và `time.After`
8. Viết chương trình đọc nhiều files concurrently và merge nội dung
9. Implement fan-out/fan-in pattern
10. Tạo một concurrent-safe counter sử dụng channels (không dùng mutex)

### Advanced Topics

1. **Context**: Viết HTTP client với timeout và cancellation sử dụng `context.Context`
2. **JSON Processing**: Parse và generate JSON, handle nested structures, custom marshaling
3. **File I/O**:
   - Đọc file lớn hiệu quả (buffered reading)
   - Viết CSV file
   - Copy file với progress tracking
4. **Testing**:
   - Viết benchmark tests cho các functions
   - Implement table-driven tests
   - Mock external dependencies
   - Use test fixtures
5. **Reflection**: Viết function generic để print struct fields và values sử dụng reflection

## OOP Concepts in Go

Trả lời bằng cách tạo file có tên là `oop-X.md` với X là số thứ tự của câu hỏi (tham khảo định dạng markdown). Bạn được phép tham khảo bất kỳ nguồn nào nhưng cố gắng chỉ viết những gì bạn thực sự hiểu. Với những gì bạn không hiểu, hãy viết dưới dạng câu hỏi. Cố gắng viết ví dụ code minh họa cho mỗi concept.

1. Go có phải là ngôn ngữ hướng đối tượng không? Giải thích cách Go implement OOP concepts
2. Struct trong Go là gì? So sánh struct trong Go với class trong Java
3. Method trong Go là gì? Phân biệt value receiver và pointer receiver
4. Interface trong Go là gì? Giải thích implicit interface implementation
5. Embedding trong Go là gì? So sánh với inheritance trong Java
6. Composition over inheritance - Go áp dụng principle này như thế nào?
7. Polymorphism trong Go được thực hiện ra sao?
8. Giải thích Empty Interface `interface{}` và type assertions
9. Type switch trong Go dùng để làm gì?
10. Giải thích về pointer trong Go và khi nào nên dùng pointer
11. Exported vs Unexported (public vs private) trong Go hoạt động như thế nào?
12. Method sets trong Go là gì? Giải thích rules của method sets với pointers và values
13. Giải thích về zero values trong Go và tại sao nó quan trọng
14. Constructor pattern trong Go - làm sao để tạo constructor cho struct?
15. Giải thích defer, panic, và recover
16. Multiple return values trong Go - best practices khi nào nên dùng
17. Giải thích về channels và goroutines trong context của concurrent programming
18. Mutex và RWMutex - khi nào nên dùng cái nào?
19. Package organization và visibility rules trong Go
20. Các Design Patterns phổ biến trong Go (Singleton, Factory, Builder, Strategy, etc.) - implement và giải thích

## Best Practices

1. **Code Organization**:

   - Một package một purpose
   - Keep functions small and focused
   - Use meaningful names

2. **Error Handling**:

   - Always handle errors, never ignore them
   - Return errors instead of panicking
   - Wrap errors with context

3. **Concurrency**:

   - Don't communicate by sharing memory; share memory by communicating
   - Use channels for communication between goroutines
   - Close channels from sender side only

4. **Testing**:

   - Write tests for all public functions
   - Use table-driven tests
   - Keep test files next to source files

5. **Documentation**:
   - Write godoc comments for exported functions
   - Include examples in documentation
   - Keep comments up-to-date

## Resources

- [A Tour of Go](https://tour.golang.org/)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

## Submission

- Create a Pull Request for each completed section
- Ensure all tests pass with good coverage
- Follow Go formatting conventions (`go fmt`)
- Run `go vet` and `golint` before submitting
- Add meaningful commit messages
