# 📘 Golang Internship 2025 – Handbook

Chào mừng các bạn đến với chương trình **Golang Internship 2025**.

Tài liệu này đóng vai trò là **kim chỉ nam (Guide)** và **quy chuẩn (Standard)** mà các bạn cần tuân thủ trong suốt quá trình làm việc. Mục tiêu không chỉ là **code chạy được**, mà là **code Clean, Maintainable và Idiomatic Go**.

---

## 🏗️ Phần 1: Coding Conventions (Quy chuẩn lập trình)

Trước khi bắt tay vào làm bài tập, các bạn **bắt buộc** phải thiết lập môi trường và tuân thủ các quy tắc sau.

### 1. Tooling & Environment

- **Go Version**: `1.21+`
- **IDE**: VS Code (khuyên dùng) hoặc GoLand
- **Formatter**:

  - Bắt buộc bật **Format On Save**
  - Sử dụng **gofumpt** (phiên bản strict hơn của `gofmt`)

- **Linter**:

  - Cài đặt `golangci-lint`
  - Chạy lệnh sau trước khi commit:

```bash
golangci-lint run ./...
```

- **Pull Request**:

  - Không được phép có **warning** từ linter

---

### 2. Naming Convention (Quy tắc đặt tên)

> Go rất khắt khe về việc đặt tên.
> **Câu thần chú**: _"Càng xa nơi khai báo, tên càng phải rõ nghĩa. Càng gần, tên càng ngắn gọn."_

#### Package

- Tên **ngắn**, toàn chữ **thường**, danh từ **số ít**
- Không dùng underscore hay camelCase

```go
// ✅ Good
package user

// ❌ Bad
package UserInfo
package user_service
```

#### Variable / Constant

- `camelCase` cho biến **private** (unexported)
- `PascalCase` cho biến **public** (exported)
- Viết tắt phải **đồng nhất**:

  - `ServeHTTP` (tốt)
  - `ServeHttp` (xấu)

- Biến vòng lặp nên ngắn gọn: `i`, `v`, `k`

#### Interface

- Thường kết thúc bằng `er`

```go
Reader
Writer
Formatter
```

---

### 3. Error Handling Style

- ❌ **KHÔNG dùng `panic`**, trừ khi lỗi nghiêm trọng khi khởi tạo chương trình (`main`)
- Luôn kiểm tra lỗi **ngay lập tức**

```go
// Good
f, err := os.Open("file.txt")
if err != nil {
    return err
}
// Logic tiếp theo ở indent ngoài cùng
```

- Wrap error để dễ trace (dùng `%w`)

```go
return fmt.Errorf("failed to open config: %w", err)
```

---

### 4. Project Layout

Áp dụng **flat structure** cho bài tập và **module structure** cho dự án lớn.

```text
.
├── cmd/                # Entry points (main.go)
├── internal/           # Private application code (project lớn)
├── pkg/                # Library code dùng bên ngoài
├── exercises/          # Code bài tập
│   ├── basic/
│   ├── string/
│   └── ...
├── go.mod
├── go.sum
└── README.md
```

---

## 🚀 Phần 2: Internship Exercises & Roadmap

### 1. Setup Project

#### Clone & Init

```bash
git clone <your-repo-url>
cd golang-internship-2025
go mod init github.com/yourusername/golang-internship-2025
```

#### Setup CI/CD

- Copy file `.github/workflows/ci.yml`
- Đăng ký **Codecov**, lấy `CODECOV_TOKEN`
- Vào **GitHub Repo Settings → Secrets → Actions**
- Thêm secret:

```text
CODECOV_TOKEN
```

#### Badges

- Cập nhật `README.md` để hiển thị:

  - Build Status
  - Code Coverage

---

### 2. Testing Requirements (Yêu cầu kiểm thử)

- **Coverage**: Target **> 80%**
- Code không có test ⇒ **chưa hoàn thành**
- **Table Driven Test**: _Bắt buộc_
- **Benchmark** (khuyến khích với thuật toán):

```go
func BenchmarkXxx(b *testing.B)
```

---

### 3. Danh sách bài tập (Exercises)

#### 🟢 Basic Algorithms

- GCD & LCM
- Sum Digits
- Prime Factorization (VD: `600 = 2^3 * 3 * 5^2`)
- Fibonacci (≤ n)
- Factorial (Recursive & Iterative)
- Month Converter (switch/map)
- Reverse Binary
- Roman Numerals (Decimal ↔ Roman)

#### 🔵 String Processing

- Palindrome
- Sum in String
- Pattern Repeat
- Run-length Encoding
- Run-length Decoding
- EAN-13 Validator
- Char Frequency (Map)

#### 🟡 Slice & Array Operations

- Sum Even / Odd
- Binary Search
- In-place Reverse
- Merge Sorted Slices (O(n))
- Find Missing Number (1..n)
- K-th Largest (QuickSelect)
- Rotate Slice

#### 🟠 Structs, Maps & Methods

- Word Count
- Person Struct (Validate + Method)
- Student Management (Embedding, GPA)
- Phonebook (CRUD – in-memory)
- Simple Cache (Map + TTL)

#### 🟣 Interfaces & Polymorphism

- Geometry (Shape: Area, Perimeter)
- Custom Sort (Sortable)
- Payment Gateway (Strategy Pattern)
- Storage System (File / Memory)
- Compression (Zip / Rar / Gzip)

#### 🔴 Error Handling & Concurrency (Advanced)

- Safe File Reader
- Banking Errors (Custom error)
- Retry Mechanism (Exponential Backoff)
- Concurrent Sum (Goroutines)
- Worker Pool
- Rate Limiter (Channel / Ticker)
- Timeout (`context.WithTimeout`, `select`)

---

## 🧠 Phần 3: OOP & Theory Assessment

Tạo thư mục `docs/` và trả lời các câu hỏi sau bằng Markdown.

**File**: `docs/oop-concepts.md`

- Go có phải ngôn ngữ OOP?
- Composition vs Inheritance
- Pointer vs Value Receiver
- Interface & Duck Typing
- Zero Values
- `nil` slice vs `nil` map
- Cơ chế `defer` (LIFO)
- `make` vs `new`
- Goroutine vs System Thread (M:N)
- Race Condition & cách detect

---

## ✅ Submission Process (Quy trình nộp bài)

### Commit Message

Tuân thủ **Conventional Commits**:

```text
feat: add fibonacci solution
test: add benchmark for prime factors
docs: update readme
```

### Pull Request

- Tạo PR vào nhánh `main`
- Đảm bảo CI (Build, Test, Lint) **xanh**
- Self-review trước khi request review
- Tag mentor

---

## 💡 Lời khuyên từ SA

> "Code là để con người đọc, chỉ tình cờ là máy tính thực thi được."

Trong Go, chúng ta ưu tiên **Clarity** hơn **Cleverness**.

- Tránh one-liner phức tạp
- Viết code để đồng nghiệp đọc vào là hiểu ngay logic

---

**Chúc các bạn có một kỳ thực tập hiệu quả và học được nhiều điều thú vị với Go!**

# 📘 Golang Internship 2025 – Handbook

Chào mừng các bạn đến với chương trình **Golang Internship 2025**.

Tài liệu này đóng vai trò là **kim chỉ nam (Guide)** và **quy chuẩn (Standard)** mà các bạn cần tuân thủ trong suốt quá trình làm việc. Mục tiêu không chỉ là **code chạy được**, mà là **code Clean, Maintainable và Idiomatic Go**.

---

## 🏗️ Phần 1: Coding Conventions (Quy chuẩn lập trình)

Trước khi bắt tay vào làm bài tập, các bạn **bắt buộc** phải thiết lập môi trường và tuân thủ các quy tắc sau.

### 1. Tooling & Environment

- **Go Version**: `1.21+`
- **IDE**: VS Code (khuyên dùng) hoặc GoLand
- **Formatter**:

  - Bắt buộc bật **Format On Save**
  - Sử dụng **gofumpt** (phiên bản strict hơn của `gofmt`)

- **Linter**:

  - Cài đặt `golangci-lint`
  - Chạy lệnh sau trước khi commit:

```bash
golangci-lint run ./...
```

- **Pull Request**:

  - Không được phép có **warning** từ linter

---

### 2. Naming Convention (Quy tắc đặt tên)

> Go rất khắt khe về việc đặt tên.
> **Câu thần chú**: _"Càng xa nơi khai báo, tên càng phải rõ nghĩa. Càng gần, tên càng ngắn gọn."_

#### Package

- Tên **ngắn**, toàn chữ **thường**, danh từ **số ít**
- Không dùng underscore hay camelCase

```go
// ✅ Good
package user

// ❌ Bad
package UserInfo
package user_service
```

#### Variable / Constant

- `camelCase` cho biến **private** (unexported)
- `PascalCase` cho biến **public** (exported)
- Viết tắt phải **đồng nhất**:

  - `ServeHTTP` (tốt)
  - `ServeHttp` (xấu)

- Biến vòng lặp nên ngắn gọn: `i`, `v`, `k`

#### Interface

- Thường kết thúc bằng `er`

```go
Reader
Writer
Formatter
```

---

### 3. Error Handling Style

- ❌ **KHÔNG dùng `panic`**, trừ khi lỗi nghiêm trọng khi khởi tạo chương trình (`main`)
- Luôn kiểm tra lỗi **ngay lập tức**

```go
// Good
f, err := os.Open("file.txt")
if err != nil {
    return err
}
// Logic tiếp theo ở indent ngoài cùng
```

- Wrap error để dễ trace (dùng `%w`)

```go
return fmt.Errorf("failed to open config: %w", err)
```

---

### 4. Project Layout

Áp dụng **flat structure** cho bài tập và **module structure** cho dự án lớn.

```text
.
├── cmd/                # Entry points (main.go)
├── internal/           # Private application code (project lớn)
├── pkg/                # Library code dùng bên ngoài
├── exercises/          # Code bài tập
│   ├── basic/
│   ├── string/
│   └── ...
├── go.mod
├── go.sum
└── README.md
```

---

## 🚀 Phần 2: Internship Exercises & Roadmap

### 1. Setup Project

#### Clone & Init

```bash
git clone <your-repo-url>
cd golang-internship-2025
go mod init github.com/yourusername/golang-internship-2025
```

#### Setup CI/CD

- Copy file `.github/workflows/ci.yml`
- Đăng ký **Codecov**, lấy `CODECOV_TOKEN`
- Vào **GitHub Repo Settings → Secrets → Actions**
- Thêm secret:

```text
CODECOV_TOKEN
```

#### Badges

- Cập nhật `README.md` để hiển thị:

  - Build Status
  - Code Coverage

---

### 2. Testing Requirements (Yêu cầu kiểm thử)

- **Coverage**: Target **> 80%**
- Code không có test ⇒ **chưa hoàn thành**
- **Table Driven Test**: _Bắt buộc_
- **Benchmark** (khuyến khích với thuật toán):

```go
func BenchmarkXxx(b *testing.B)
```

---

### 3. Danh sách bài tập (Exercises)

#### 🟢 Basic Algorithms

- GCD & LCM
- Sum Digits
- Prime Factorization (VD: `600 = 2^3 * 3 * 5^2`)
- Fibonacci (≤ n)
- Factorial (Recursive & Iterative)
- Month Converter (switch/map)
- Reverse Binary
- Roman Numerals (Decimal ↔ Roman)

#### 🔵 String Processing

- Palindrome
- Sum in String
- Pattern Repeat
- Run-length Encoding
- Run-length Decoding
- EAN-13 Validator
- Char Frequency (Map)

#### 🟡 Slice & Array Operations

- Sum Even / Odd
- Binary Search
- In-place Reverse
- Merge Sorted Slices (O(n))
- Find Missing Number (1..n)
- K-th Largest (QuickSelect)
- Rotate Slice

#### 🟠 Structs, Maps & Methods

- Word Count
- Person Struct (Validate + Method)
- Student Management (Embedding, GPA)
- Phonebook (CRUD – in-memory)
- Simple Cache (Map + TTL)

#### 🟣 Interfaces & Polymorphism

- Geometry (Shape: Area, Perimeter)
- Custom Sort (Sortable)
- Payment Gateway (Strategy Pattern)
- Storage System (File / Memory)
- Compression (Zip / Rar / Gzip)

#### 🔴 Error Handling & Concurrency (Advanced)

- Safe File Reader
- Banking Errors (Custom error)
- Retry Mechanism (Exponential Backoff)
- Concurrent Sum (Goroutines)
- Worker Pool
- Rate Limiter (Channel / Ticker)
- Timeout (`context.WithTimeout`, `select`)

---

## 🧠 Phần 3: OOP & Theory Assessment

Tạo thư mục `docs/` và trả lời các câu hỏi sau bằng Markdown.

**File**: `docs/oop-concepts.md`

- Go có phải ngôn ngữ OOP?
- Composition vs Inheritance
- Pointer vs Value Receiver
- Interface & Duck Typing
- Zero Values
- `nil` slice vs `nil` map
- Cơ chế `defer` (LIFO)
- `make` vs `new`
- Goroutine vs System Thread (M:N)
- Race Condition & cách detect

---

## ✅ Submission Process (Quy trình nộp bài)

### Commit Message

Tuân thủ **Conventional Commits**:

```text
feat: add fibonacci solution
test: add benchmark for prime factors
docs: update readme
```

### Pull Request

- Tạo PR vào nhánh `main`
- Đảm bảo CI (Build, Test, Lint) **xanh**
- Self-review trước khi request review
- Tag mentor

---

## 💡 Lời khuyên từ SA

> "Code là để con người đọc, chỉ tình cờ là máy tính thực thi được."

Trong Go, chúng ta ưu tiên **Clarity** hơn **Cleverness**.

- Tránh one-liner phức tạp
- Viết code để đồng nghiệp đọc vào là hiểu ngay logic

---

**Chúc các bạn có một kỳ thực tập hiệu quả và học được nhiều điều thú vị với Go!**
