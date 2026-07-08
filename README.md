# 🚀 Go Mastery Roadmap (বাংলা)

[![GitHub](https://img.shields.io/badge/GitHub-shagorrobidas-181717?style=flat-square&logo=github)](https://github.com/shagorrobidas/learn-go-programming)
[![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)](LICENSE)
[![Language](https://img.shields.io/badge/Language-Bangla-blue?style=flat-square)](README.md)
[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![Status](https://img.shields.io/badge/Status-Active-brightgreen?style=flat-square)](README.md)

> **Go (Golang) শেখার সম্পূর্ণ রোডম্যাপ - শুরু থেকে Advanced পর্যন্ত। এই রিপোজিটরিতে আপনি পাবেন বাংলায় লেখা সম্পূর্ণ গাইড, প্র্যাক্টিক্যাল উদাহরণ এবং রিয়েল-ওয়ার্ল্ড প্রজেক্ট।**

---

## 📋 বিষয়বস্তু (Table of Contents)

- [পরিচিতি](#পরিচিতি)
- [এই Repository-তে কী পাবেন](#এই-repository-তে-কী-পাবেন)
- [Learning Roadmap](#learning-roadmap)
- [Go Basics](#go-basics)
- [Functions](#functions)
- [Control Flow](#control-flow)
- [Collections](#collections)
- [Pointers](#pointers)
- [Struct এবং Methods](#struct-এবং-methods)
- [Interface](#interface)
- [Packages এবং Modules](#packages-এবং-modules)
- [Error Handling](#error-handling)
- [Concurrency](#concurrency)
- [Advanced Topics](#advanced-topics)
- [প্রজেক্ট এবং চ্যালেঞ্জ](#প্রজেক্ট-এবং-চ্যালেঞ্জ)
- [Interview Preparation](#interview-preparation)
- [Study Schedule](#study-schedule)
- [Resources](#resources)
- [Contributing](#contributing)
- [License](#license)

---

## 🎯 পরিচিতি

### Go কী?

**Go** (যা **Golang** নামেও পরিচিত) একটি আধুনিক, সুদ্রঢ় এবং দক্ষ প্রোগ্রামিং ভাষা। এটি **2007 সালে Google-এ তৈরি** করা হয় এবং **2009 সালে সর্বজনীনভাবে প্রকাশ করা হয়**। Go একটি **compiled language** যা **Unix/Linux আধারভিত্তিক সিস্টেমের জন্য অপ্টিমাইজ করা** হয়েছে।

Go-এর স্রষ্টারা হলেন:
- **Robert Griesemer**
- **Rob Pike**
- **Ken Thompson**

### Golang কেন তৈরি করা হয়েছে?

Google তাদের বিশাল সার্ভার অবকাঠামো পরিচালনার জন্য একটি নতুন ভাষার প্রয়োজন অনুভব করেছিল। তারা এমন একটি ভাষা চেয়েছিল যা:

- **দ্রুত কম্পাইল হয়**
- **মেশিন কোডে দ্রুত চলে**
- **Concurrency-তে দক্ষ**
- **সহজ এবং সাধারণবোধগম্য সিনট্যাক্স**
- **একটি একক executable বাইনারি তৈরি করে**

### Go কারা ব্যবহার করে?

বিশ্বের শীর্ষস্থানীয় কোম্পানিগুলি Go ব্যবহার করে:

| কোম্পানি | ব্যবহার |
|---------|--------|
| **Google** | ক্ল���উড প্ল্যাটফর্ম, সেবা |
| **Facebook** | ডেটা প্রসেসিং |
| **Uber** | মাইক্রোসার্ভিসেস |
| **Netflix** | পারফরম্যান্স টুলস |
| **Dropbox** | ব্যাকএন্ড সিস্টেম |
| **Docker** | কন্টেইনারাইজেশন |
| **Kubernetes** | অর্কেস্ট্রেশন |
| **MongoDB** | ডাটাবেস ড্রাইভার |
| **HashiCorp** | ক্লাউড ইনফ্রাস্ট্রাকচার টুলস |
| **Slack** | ব্যাকএন্ড সিস্টেম |

### Go কোথায় ব্যবহার হয়?

Go-এর প্রধান ব্যবহার ক্ষেত্র:

✅ **Backend Web Development** - REST API, Microservices
✅ **Cloud & DevOps** - Docker, Kubernetes, Terraform
✅ **System Programming** - অপারেটিং সিস্টেম, টুলস
✅ **Networking** - HTTP সার্ভার, gRPC
✅ **Data Processing** - বড় ডাটা প্রসেসিং
✅ **CLI Tools** - কমান্ড লাইন অ্যাপ্লিকেশন
✅ **Real-time Systems** - গেমিং, চ্যাট অ্যাপ্লিকেশ���

### কেন Go শিখবেন?

#### ১. **দ্রুত পারফরম্যান্স**
Go সরাসরি মেশিন কোডে কম্পাইল হয়, যা Python বা JavaScript-এর তুলনায় **১০-১০০ গুণ দ্রুত**।

#### २. **Concurrency সহজ**
Go-এর **Goroutines** এবং **Channels** ব্যবহার করে হাজার হাজার concurrent operations সহজেই পরিচালনা করা যায়।

#### ३. **একটি Executable Binary**
একটি স্ট্যান্ডঅ্যালোন executable তৈরি হয় যা কোনো dependencies ছাড়াই চলে।

#### ४. **শিখতে সহজ**
Go-এর সিনট্যাক্স সহজ এবং পরিষ্কার। **"There should be one way to do things"** - এই নীতি অনুসরণ করে।

#### ५. **উৎপাদনশীলতা**
দ্রুত উন্নয়ন, শক্তিশালী standard library, এবং কম boilerplate কোড।

#### ६. **চাকরির সুযোগ**
Go ডেভেলপারদের জন্য শীর্ষ প্রযুক্তি কোম্পানিতে উচ্চ চাহিদা রয়েছে।

### Go-এর প্রধান সুবিধা

```
┌─────────────────────────────────────────┐
│          Go-এর সুবিধা                   │
├─────────────────────────────────────────┤
│ 1. দ্রুত কম্পাইলেশন                      │
│ 2. Concurrency নেটিভ সাপোর্ট           │
│ 3. সিম্পল এবং পরিষ্কার সিনট্যাক্স      │
│ 4. Strong Standard Library             │
│ 5. Cross-platform সাপোর্ট             │
│ 6. Static Typing + সুবিধা              │
│ 7. বিল্ট-ইন টেস্টিং ফ্রেমওয়ার্ক        │
│ 8. ডকুমেন্টেশন টুলস অন্তর্ভুক্ত       │
└─────────────────────────────────────────┘
```

### Go-এর ভবিষ্যৎ

Go ক্রমাগত বৃদ্ধি পাচ্ছে এবং নতুন বৈশিষ্ট্য যোগ হচ্ছে:
- **Go 1.21**: Generics এবং improved performance
- **ভবিষ্যৎ**: আরও better error handling এবং optimization

---

## 📦 এই Repository-তে কী পাবেন?

এই সম্পূর্ণ গাইডে আপনি পাবেন:

### ✅ **Go Fundamentals**
- Installation এবং Setup
- Workspace Configuration
- প্রথম Program লেখা
- Variables এবং Constants
- Data Types এবং Operations

### ✅ **Intermediate Concepts**
- Functions এবং Higher-order Functions
- Pointers এবং Memory Management
- Structs এবং Methods
- Interfaces এবং Polymorphism
- Error Handling এবং Panic/Recover

### ✅ **Advanced Topics**
- Goroutines এবং Concurrency Patterns
- Channels এবং Channel Operations
- Context এবং Cancellation
- Generics এবং Type Parameters
- Memory Optimization এবং Profiling

### ✅ **Practical Applications**
- REST API Development
- HTTP Client এবং Server
- Database Integration (PostgreSQL, MySQL)
- JSON Processing
- File এবং Directory Operations

### ✅ **Frameworks এবং Tools**
- Gin Web Framework
- Fiber Framework
- GORM ORM Library
- Redis Integration
- WebSocket Implementation

### ✅ **Enterprise Patterns**
- Clean Architecture
- Design Patterns
- Microservices
- Docker এবং Containerization
- Kubernetes Orchestration

### ✅ **Project-based Learning**
- Beginner Projects
- Intermediate Projects
- Advanced Projects
- Real-world Scenarios

### ✅ **Interview Preparation**
- ১০০+ সাক্ষাৎকার প্রশ্ন
- বিস্তারিত উত্তর
- Coding কিলার প্রশ্ন
- System Design প্রশ্ন

---

## 🗺️ Learning Roadmap

| সপ্তাহ | টপিক | কঠিনতা | সময় | প্র্যাক্টিস | স্ট্যাটাস |
|------|------|--------|------|-----------|--------|
| Week 1 | Setup এবং Hello World | ⭐ | 2-3 ঘন্টা | Basic Programs | ✅ |
| Week 1 | Variables এবং Constants | ⭐ | 3-4 ঘন্টা | 10+ Programs | ✅ |
| Week 2 | Data Types এবং Operators | ⭐ | 4-5 ঘন্টা | Type Conversion | ✅ |
| Week 2 | Input/Output এবং Formatting | ⭐ | 3-4 ঘন্টা | String Manipulation | ✅ |
| Week 3 | If-Else এবং Conditionals | ⭐ | 3-4 ঘন্টা | 15+ Problems | ✅ |
| Week 3 | Loops এবং Range | ⭐ | 4-5 ঘন্টা | Pattern Programs | ✅ |
| Week 4 | Functions এবং Parameters | ⭐⭐ | 5-6 ঘন্টা | 20+ Functions | ✅ |
| Week 4 | Multiple Return এবং Named Returns | ⭐⭐ | 4-5 ঘন্টা | Practice Problems | ✅ |
| Week 5 | Arrays এবং Slices | ⭐⭐ | 5-6 ঘন্টা | Slice Operations | ✅ |
| Week 5 | Maps এবং Key-Value Operations | ⭐⭐ | 4-5 ঘন্টা | Real-world Examples | ✅ |
| Week 6 | Strings এবং Runes | ⭐⭐ | 4-5 ঘন্টা | String Processing | ✅ |
| Week 6 | Pointers এবং Memory | ⭐⭐ | 5-6 ঘন্টা | Pointer Exercises | ✅ |
| Week 7 | Structs এবং Methods | ⭐⭐ | 6-7 ঘন্টা | OOP Concepts | ✅ |
| Week 7 | Embedding এবং Composition | ⭐⭐⭐ | 5-6 ঘন্টা | Design Patterns | ✅ |
| Week 8 | Interfaces | ⭐⭐⭐ | 6-7 ঘন্টা | Polymorphism | ✅ |
| Week 8 | Type Assertion এবং Type Switch | ⭐⭐⭐ | 4-5 ঘন্টা | Advanced Type Operations | ✅ |
| Week 9 | Packages এবং Imports | ⭐⭐ | 4-5 ঘন্টা | Module Organization | ✅ |
| Week 9 | Error Handling Best Practices | ⭐⭐⭐ | 5-6 ঘন্টা | Error Wrapping | ✅ |
| Week 10 | Panic এবং Recover | ⭐⭐⭐ | 4-5 ঘন্টা | Exception Handling | ✅ |
| Week 10 | Defer Statement | ⭐⭐ | 3-4 ঘন্টা | Resource Management | ✅ |
| Week 11 | JSON Processing | ⭐⭐ | 5-6 ঘন্টা | Struct Tags | ✅ |
| Week 11 | File Operations | ⭐⭐ | 5-6 ঘন্টা | File I/O | ✅ |
| Week 12 | HTTP Client এবং Server | ⭐⭐⭐ | 6-7 ঘন্টা | REST API | ✅ |
| Week 12 | Middleware এবং Routing | ⭐⭐⭐ | 5-6 ঘন্টা | Web Framework | ✅ |
| Week 13 | Goroutines | ⭐⭐⭐ | 6-7 ঘন্টা | Concurrency | ✅ |
| Week 13 | Channels এবং Communication | ⭐⭐⭐⭐ | 7-8 ঘন্টা | Channel Patterns | ✅ |
| Week 14 | Select এবং Multiplexing | ⭐⭐⭐⭐ | 5-6 ঘন্টা | Complex Patterns | ✅ |
| Week 14 | Context এবং Cancellation | ⭐⭐⭐ | 5-6 ঘন্টা | Timeout Handling | ✅ |
| Week 15 | Mutex এবং Synchronization | ⭐⭐⭐⭐ | 6-7 ঘন্টা | Thread Safety | ✅ |
| Week 15-30 | Advanced Topics এবং Projects | ⭐⭐⭐⭐⭐ | 60+ ঘন্টা | Real Applications | 🔄 |

---

## 🚀 Go Basics

### ১. Installation এবং Setup

#### ✔️ এটি কী?

Go ডেভেলপমেন্ট শুরু করার জন্য আমাদের প্রথমে Go ইনস্টল করতে হবে এবং পরিবেশ সেটআপ করতে হবে।

#### ✔️ কীভাবে ইনস্টল করবেন?

**Linux/Mac-এ:**
```bash
# Homebrew ব্যবহার করে (Mac)
brew install go

# Linux-এ (Ubuntu/Debian)
sudo apt-get install golang-go

# উভয় ক্ষেত্রে সংস্করণ চেক করুন
go version
```

**Windows-এ:**
- [golang.org](https://golang.org/dl) থেকে installer ডাউনলোড করুন
- Installer চালান এবং ডিফল্ট পথে ইনস্টল করুন
- `go version` চেক করুন Command Prompt-এ

#### ✔️ GOPATH এবং GOROOT

```bash
# GOROOT - Go ইনস্টলেশন ডিরেক্টরি
echo $GOROOT

# GOPATH - আমাদের কর্মক্ষেত্র
echo $GOPATH

# GOPATH সেট করুন
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

#### ✔️ প্রথম Program: Hello World

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World! 👋 Go-এ স্বাগতম!")
}
```

**Output:**
```
Hello, World! 👋 Go-এ স্বাগতম!
```

**চালানোর উপায়:**
```bash
# সরাসরি চালান
go run main.go

# কম্পাইল করে executable তৈরি করুন
go build main.go
./main
```

#### ✔️ IDE সেটআপ

**VS Code-এ:**
1. VS Code ইনস্টল করুন
2. "Go" এক্সটেনশন ইনস্টল করুন (Golang)
3. `Ctrl+Shift+P` → "Go: Install/Update Tools" চলান

**GoLand (JetBrains):**
- Professional IDE তবে Paid
- সম্পূর্ণ Go সাপোর্ট এবং ডিবাগিং

---

### २. Variables এবং Constants

#### ✔️ Variables কী?

Variable হল মূল্য সংরক্ষণের জন্য নামকৃত স্টোরেজ স্থান।

```go
package main

import "fmt"

func main() {
    // Method 1: var keyword
    var name string = "আল করিম"
    var age int = 25
    var height float64 = 5.9

    // Method 2: short declaration (শুধুমাত্র function-এ)
    city := "ঢাকা"
    score := 95

    // Multiple declarations
    var (
        firstname = "মুহাম্মদ"
        lastname = "করিম"
        email = "karim@example.com"
    )

    fmt.Println("নাম:", name)
    fmt.Println("বয়স:", age)
    fmt.Println("শহর:", city)
}
```

**Output:**
```
নাম: আল করিম
বয়স: 25
শহর: ঢাকা
```

#### ✔️ Constants কী?

Constant হল এমন মূল্য যা একবার সেট করার পর পরিবর্তন করা যায় না।

```go
package main

import "fmt"

func main() {
    // Constants
    const (
        MaxPlayers = 100
        AppName = "আমাদের গেম"
        Version = "1.0"
        Pi = 3.14159
    )

    fmt.Println("অ্যাপ নাম:", AppName)
    fmt.Println("সংস্করণ:", Version)
    fmt.Println("সর্বোচ্চ খেলোয়াড়:", MaxPlayers)

    // এটি ত্রুটি দেবে:
    // MaxPlayers = 200 // ❌ Error
}
```

**সাধারণ ভুল:**

```go
// ❌ ভুল
const name string = "করিম"

// ✅ সঠিক
const name = "করিম"
```

---

### ३. Data Types

#### ✔️ প্রধান Data Types

```go
package main

import "fmt"

func main() {
    // Numeric Types
    var intNum int = 42
    var floatNum float64 = 3.14
    var uint8Num uint8 = 255

    // String
    var text string = "বাংলায় স্বাগতম"

    // Boolean
    var isActive bool = true

    // Rune এবং Byte
    var letter rune = 'ক'  // Unicode character
    var ascii byte = 'A'   // ASCII character

    fmt.Printf("Integer: %d (Type: %T)\n", intNum, intNum)
    fmt.Printf("Float: %.2f (Type: %T)\n", floatNum, floatNum)
    fmt.Printf("String: %s (Type: %T)\n", text, text)
    fmt.Printf("Boolean: %v (Type: %T)\n", isActive, isActive)
    fmt.Printf("Rune: %c (Type: %T)\n", letter, letter)
}
```

**Output:**
```
Integer: 42 (Type: int)
Float: 3.14 (Type: float64)
String: বাংলায় স্বাগতম (Type: string)
Boolean: true (Type: bool)
Rune: ক (Type: int32)
```

#### ✔️ Type Conversion

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    // String থেকে Int
    var str = "123"
    intVal, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println("রূপান্তর ব্যর্থ:", err)
    }
    fmt.Printf("String '%s' → Integer %d\n", str, intVal)

    // Int থেকে String
    var num = 456
    strNum := strconv.Itoa(num)
    fmt.Printf("Integer %d → String '%s'\n", num, strNum)

    // Float থেকে Int
    var f float64 = 7.89
    i := int(f)
    fmt.Printf("Float %.2f → Integer %d\n", f, i)
}
```

**Interview প্রশ্ন:**

❓ **Q: int এবং int64-এর মধ্যে পার্থক্য কী?**
> A: `int` প্ল্যাটফর্ম-নির্ভর (32 বা 64 বিট), যখন `int64` সবসময় 64 বিট।

❓ **Q: নিম্নোক্ত কোড কেন ত্রুটি দেয়?**
```go
var x int = 10
var y float64 = x + 5.5 // ❌ Error
```
> A: Go-তে implicit type conversion নেই। স্পষ্ট রূপান্তর প্রয়োজন: `var y float64 = float64(x) + 5.5`

---

### ४. Operators

#### ✔️ Arithmetic Operators

```go
package main

import "fmt"

func main() {
    a := 10
    b := 3

    fmt.Printf("%d + %d = %d\n", a, b, a+b)      // যোগ
    fmt.Printf("%d - %d = %d\n", a, b, a-b)      // বিয়োগ
    fmt.Printf("%d × %d = %d\n", a, b, a*b)      // গুণ
    fmt.Printf("%d ÷ %d = %d\n", a, b, a/b)      // ভাগ
    fmt.Printf("%d %% %d = %d\n", a, b, a%b)     // অবশিষ্ট
}
```

#### ✔️ Comparison এবং Logical Operators

```go
package main

import "fmt"

func main() {
    age := 25
    income := 50000

    // Comparison
    fmt.Println("age > 18:", age > 18)
    fmt.Println("age == 25:", age == 25)
    fmt.Println("income <= 50000:", income <= 50000)

    // Logical
    fmt.Println("age > 18 && income > 30000:", age > 18 && income > 30000)
    fmt.Println("age < 18 || income > 40000:", age < 18 || income > 40000)
    fmt.Println("!(age < 18):", !(age < 18))
}
```

---

## 🔧 Functions

### Functions এর মূল ধারণা

#### ✔️ Function কী?

Function হল একটি নির্দিষ্ট কাজ সম্পাদনের জন্য কোডের পুনঃব্যবহারযোগ্য ব্লক।

```go
package main

import "fmt"

// সাধারণ function
func greet(name string) {
    fmt.Printf("আল-সালাম, %s!\n", name)
}

// Return value সহ function
func add(a, b int) int {
    return a + b
}

// Multiple return values
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("শূন্য দিয়ে ভাগ করা যায় না")
    }
    return a / b, nil
}

func main() {
    greet("করিম")
    
    result := add(10, 20)
    fmt.Println("যোগফল:", result)

    quotient, err := divide(20, 4)
    if err != nil {
        fmt.Println("ত্রুটি:", err)
    } else {
        fmt.Println("ভাগফল:", quotient)
    }
}
```

**Output:**
```
আল-সালাম, করিম!
যোগফল: 30
ভাগফল: 5
```

#### ✔️ Named Return Values

```go
package main

import "fmt"

func calculateBMI(weight, height float64) (bmi float64, status string) {
    bmi = weight / (height * height)
    
    if bmi < 18.5 {
        status = "কম ওজন"
    } else if bmi < 25 {
        status = "স্বাভাবিক"
    } else {
        status = "অতিরিক্ত ওজন"
    }
    
    return  // স্বয়ংক্রিয়ভাবে bmi এবং status রিটার্ন করে
}

func main() {
    bmi, status := calculateBMI(70, 1.75)
    fmt.Printf("BMI: %.2f - Status: %s\n", bmi, status)
}
```

#### ✔️ Variadic Functions

```go
package main

import "fmt"

func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    fmt.Println("যোগ:", sum(1, 2, 3, 4, 5))           // 15
    fmt.Println("যোগ:", sum(10, 20, 30))              // 60
    fmt.Println("যোগ:", sum())                        // 0

    // Slice spread করা
    nums := []int{100, 200, 300}
    fmt.Println("যোগ:", sum(nums...))                 // 600
}
```

#### ✔️ Anonymous Functions এবং Closures

```go
package main

import "fmt"

func main() {
    // Anonymous function
    multiply := func(a, b int) int {
        return a * b
    }
    fmt.Println("গুণফল:", multiply(6, 7))

    // Closure - বাহ্যিক variable অ্যাক্সেস করে
    counter := 0
    increment := func() int {
        counter++
        return counter
    }
    
    fmt.Println(increment())  // 1
    fmt.Println(increment())  // 2
    fmt.Println(increment())  // 3
}
```

#### ✔️ Higher-Order Functions

```go
package main

import "fmt"

// Function যা function গ্রহণ করে
func applyOperation(a, b int, operation func(int, int) int) int {
    return operation(a, b)
}

func main() {
    add := func(x, y int) int { return x + y }
    subtract := func(x, y int) int { return x - y }

    fmt.Println("যোগ:", applyOperation(10, 5, add))
    fmt.Println("বিয়োগ:", applyOperation(10, 5, subtract))
}
```

#### ✔️ Common Mistakes

```go
// ❌ ভুল - আর্গুমেন্ট ধরন মিলছে না
func add(a int, b int) int {
    return a + b
}
result := add(5, "10")  // Error!

// ✅ সঠিক - সমস্ত আর্গুমেন্ট সঠিক ধরনের
result := add(5, 10)

// ❌ ভুল - Return value ব্যবহার করছি না
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("error")
    }
    return a / b, nil
}
divide(10, 0)  // ত্রুটি উপেক্ষা করা হয়েছে

// ✅ সঠিক - ত্রুটি পরীক্ষা করা
result, err := divide(10, 0)
if err != nil {
    fmt.Println("ত্রুটি:", err)
}
```

---

## 🎛️ Control Flow

### Conditional Statements

#### ✔️ If-Else Statement

```go
package main

import "fmt"

func gradeStudent(score int) {
    if score >= 80 {
        fmt.Println("গ্রেড: A")
    } else if score >= 70 {
        fmt.Println("গ্রেড: B")
    } else if score >= 60 {
        fmt.Println("গ্রেড: C")
    } else {
        fmt.Println("ফেল")
    }
}

func main() {
    gradeStudent(85)  // গ্রেড: A
    gradeStudent(72)  // গ্রেড: B
}
```

#### ✔️ Switch Statement

```go
package main

import "fmt"

func describeDay(day int) {
    switch day {
    case 1:
        fmt.Println("সোমবার")
    case 2:
        fmt.Println("মঙ্গলবার")
    case 3:
        fmt.Println("বুধবার")
    case 4:
        fmt.Println("বৃহস্পতিবার")
    case 5:
        fmt.Println("শুক্রবার")
    case 6:
        fmt.Println("শনিবার")
    case 7:
        fmt.Println("রবিবার")
    default:
        fmt.Println("অমান্য দিন")
    }
}

func main() {
    describeDay(3)  // বুধবার
}
```

#### ✔️ Switch with fallthrough

```go
package main

import "fmt"

func main() {
    grade := "A"
    
    switch grade {
    case "A":
        fmt.Println("চমৎকার!")
        fallthrough  // পরবর্তী case-েও যান
    case "B":
        fmt.Println("ভাল")
        fallthrough
    case "C":
        fmt.Println("সন্তোষজনক")
    default:
        fmt.Println("উন্নতির প্রয়োজন")
    }
}
```

### Loops

#### ✔️ For Loop

```go
package main

import "fmt"

func main() {
    // Traditional for loop
    for i := 1; i <= 5; i++ {
        fmt.Printf("%d ", i)
    }
    fmt.Println()  // নতুন লাইন

    // While-এর মতো loop
    count := 1
    for count <= 3 {
        fmt.Printf("গণনা: %d\n", count)
        count++
    }

    // Infinite loop
    i := 0
    for {
        if i >= 3 {
            break
        }
        fmt.Printf("অসীম loop: %d\n", i)
        i++
    }
}
```

**Output:**
```
1 2 3 4 5
গণনা: 1
গণনা: 2
গণনা: 3
অসীম loop: 0
অসীম loop: 1
অসীম loop: 2
```

#### ✔️ Range Loop

```go
package main

import "fmt"

func main() {
    // String-এ range
    text := "করিম"
    for index, char := range text {
        fmt.Printf("Index: %d, Character: %c\n", index, char)
    }

    // Slice-এ range
    numbers := []int{10, 20, 30, 40}
    for i, num := range numbers {
        fmt.Printf("সূচক: %d, মূল্য: %d\n", i, num)
    }

    // Map-এ range
    person := map[string]string{
        "নাম": "করিম",
        "শহর": "ঢাকা",
    }
    for key, value := range person {
        fmt.Printf("%s: %s\n", key, value)
    }
}
```

#### ✔️ Break এবং Continue

```go
package main

import "fmt"

func main() {
    // Break - loop থেকে বেরিয়ে যান
    for i := 1; i <= 10; i++ {
        if i == 5 {
            break  // 5-এ loop শেষ হয়
        }
        fmt.Printf("%d ", i)
    }
    fmt.Println("\nBreak loop শেষ")

    // Continue - পরবর্তী iteration-এ যান
    for i := 1; i <= 5; i++ {
        if i == 3 {
            continue  // 3 skip করুন
        }
        fmt.Printf("%d ", i)
    }
}
```

**Output:**
```
1 2 3 4
Break loop শেষ
1 2 4 5
```

#### ✔️ Labeled Loops

```go
package main

import "fmt"

func main() {
    outerLoop:
    for i := 1; i <= 3; i++ {
        for j := 1; j <= 3; j++ {
            if j == 2 {
                break outerLoop  // বাইরের loop ভাঙুন
            }
            fmt.Printf("(%d,%d) ", i, j)
        }
    }
}
```

---

## 📚 Collections

### Arrays

#### ✔️ Array কী?

Array হল একই ধরনের সংগ্রহ যার একটি নির্দিষ্ট দৈর্ঘ্য রয়েছে।

```go
package main

import "fmt"

func main() {
    // Array ঘোষণা এবং সূচনা করা
    var numbers [5]int = [5]int{10, 20, 30, 40, 50}
    
    // সংক্ষিপ্ত ঘোষণা
    fruits := [3]string{"আম", "কলা", "কমলা"}

    // Array অ্যাক্সেস করা
    fmt.Println("প্রথম সংখ্যা:", numbers[0])
    fmt.Println("ফল:", fruits)

    // Array এর দৈর্ঘ্য
    fmt.Println("দৈর্ঘ্য:", len(numbers))

    // Array সংশোধন করা
    numbers[0] = 100
    fmt.Println("সংশোধিত array:", numbers)
}
```

### Slices

#### ✔️ Slice কী?

Slice হল dynamic sized array যা array-এর মতো কাজ করে কিন্তু পরিবর্তনশীল দৈর্ঘ্যের।

```go
package main

import "fmt"

func main() {
    // Slice তৈরি করা
    var numbers []int = []int{1, 2, 3, 4, 5}
    
    // সংক্ষিপ্ত ঘোষণা
    colors := []string{"লাল", "সবুজ", "নীল"}

    // Slice indexing
    fmt.Println("প্রথম সংখ্যা:", numbers[0])
    fmt.Println("শেষ সংখ্যা:", numbers[len(numbers)-1])

    // Slicing operations
    fmt.Println("Index 1-4:", numbers[1:4])  // [2, 3, 4]
    fmt.Println("Index 2-:", numbers[2:])    // [3, 4, 5]
    fmt.Println("Index -3:", numbers[:3])    // [1, 2, 3]

    // Slice append করা
    numbers = append(numbers, 6, 7)
    fmt.Println("Append-এর পর:", numbers)

    // Slice তৈরি করা capacity সহ
    capacity := make([]int, 5, 10)  // দৈর্ঘ্য: 5, Capacity: 10
    fmt.Println("Capacity:", cap(capacity))
}
```

**Output:**
```
প্রথম সংখ্যা: 1
শেষ সংখ্যা: 5
Index 1-4: [2 3 4]
Index 2-: [3 4 5]
Index -3: [1 2 3]
Append-এর পর: [1 2 3 4 5 6 7]
Capacity: 10
```

### Maps

#### ✔️ Map কী?

Map হল unordered collection of key-value pairs।

```go
package main

import "fmt"

func main() {
    // Map তৈরি করা
    person := map[string]string{
        "নাম": "করিম",
        "শহর": "ঢাকা",
        "দেশ": "বাংলাদেশ",
    }

    // Map অ্যাক্সেস করা
    fmt.Println("নাম:", person["নাম"])
    fmt.Println("পুরো person:", person)

    // নতুন key যোগ করা
    person["পেশা"] = "ইঞ্জিনিয়ার"

    // Key-value iterate করা
    for key, value := range person {
        fmt.Printf("%s: %s\n", key, value)
    }

    // Key উপস্থিতি পরীক্ষা করা
    if value, exists := person["বয়স"]; exists {
        fmt.Println("বয়স:", value)
    } else {
        fmt.Println("বয়স key পাওয়া যায়নি")
    }

    // Key মুছে ফেলা
    delete(person, "পেশা")
    fmt.Println("delete-এর পর:", person)
}
```

### Strings এবং String Operations

#### ✔️ String Processing

```go
package main

import (
    "fmt"
    "strings"
    "unicode/utf8"
)

func main() {
    text := "আমরা বাংলাদেশী"

    // Length (bytes)
    fmt.Println("Bytes:", len(text))

    // Length (runes - characters)
    fmt.Println("Characters:", utf8.RuneCountInString(text))

    // String concatenation
    greeting := "আল-সালাম"
    name := "করিম"
    fullGreeting := greeting + " " + name
    fmt.Println(fullGreeting)

    // String functions
    fmt.Println("Contains 'বাংলা':", strings.Contains(text, "বাংলা"))
    fmt.Println("ToUpper:", strings.ToUpper(text))
    fmt.Println("Split:", strings.Split(text, " "))
    fmt.Println("Replace:", strings.ReplaceAll(text, "আমরা", "তারা"))
}
```

---

## 🔍 Pointers

### Pointer বোঝা

#### ✔️ Pointer কী?

Pointer হল একটি variable যা অন্য variable-এর memory address ধারণ করে।

```
┌──────────────┐
│  Variable    │
│  নাম: করিম   │
│  Address:    │
│  0x1234FF00  │
└──────────────┘
      ↑
      │
      └──────────────┐
                    │
              ┌──────┴─────┐
              │  Pointer   │
              │  যা address│
              │  ধারণ করে  │
              └────────────┘
```

```go
package main

import "fmt"

func main() {
    age := 25

    // & operator - address of
    ptr := &age
    
    fmt.Printf("Variable value: %d\n", age)
    fmt.Printf("Variable address: %p\n", &age)
    fmt.Printf("Pointer value: %p\n", ptr)

    // * operator - dereference
    fmt.Printf("Pointer দ্বারা value: %d\n", *ptr)

    // Pointer মাধ্যমে value পরিবর্তন করা
    *ptr = 30
    fmt.Printf("Updated age: %d\n", age)  // 30
}
```

**Output:**
```
Variable value: 25
Variable address: 0xc00008a008
Pointer value: 0xc00008a008
Pointer দ্বারা value: 25
Updated age: 30
```

#### ✔️ Pointer Functions

```go
package main

import "fmt"

// Value receiver - copy পাস করে
func incrementByValue(num int) {
    num++  // শুধুমাত্র local copy পরিবর্তন হয়
}

// Pointer receiver - reference পাস করে
func incrementByPointer(num *int) {
    *num++  // original value পরিবর্তন হয়
}

func main() {
    value := 10

    incrementByValue(value)
    fmt.Println("Value-এর পর:", value)  // 10 (পরিবর্তন নেই)

    incrementByPointer(&value)
    fmt.Println("Pointer-এর পর:", value)  // 11 (পরিবর্তন হয়েছে)
}
```

#### ✔️ Common Mistakes

```go
// ❌ ভুল - nil pointer dereference
var ptr *int  // nil
value := *ptr  // PANIC!

// ✅ সঠিক - nil check করুন
if ptr != nil {
    value := *ptr
}

// ❌ ভুল - address-এর address
var x int = 5
ptr := &&x  // Error

// ✅ সঠিক
ptr := &x  // Single level pointer
```

---

## 🏗️ Struct এবং Methods

### Struct বোঝা

#### ✔️ Struct কী?

Struct হল একই রকমের data fields একসাথে সংগ্রহ করার উপায়।

```go
package main

import "fmt"

// Struct definition
type Person struct {
    Name    string
    Age     int
    City    string
    Salary  float64
}

func main() {
    // Struct initialization
    person := Person{
        Name:   "করিম",
        Age:    28,
        City:   "ঢাকা",
        Salary: 50000,
    }

    fmt.Println("নাম:", person.Name)
    fmt.Println("বয়স:", person.Age)
    fmt.Println("শহর:", person.City)
    fmt.Println("বেতন:", person.Salary)

    // Field পরিবর্তন করা
    person.Age = 29
    fmt.Println("আপডেট বয়স:", person.Age)
}
```

#### ✔️ Methods

```go
package main

import "fmt"

type Rectangle struct {
    Width  float64
    Height float64
}

// Method - value receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Method - pointer receiver (modification-এর জন্য)
func (r *Rectangle) ScaleUp(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

func main() {
    rect := Rectangle{Width: 5, Height: 10}

    fmt.Printf("আসল ক্ষেত্রফল: %.2f\n", rect.Area())  // 50

    rect.ScaleUp(2)
    fmt.Printf("স্কেলিং-এর পর: %.2f\n", rect.Area())   // 200
}
```

#### ✔️ Embedding এবং Composition

```go
package main

import "fmt"

type Address struct {
    Street string
    City   string
    Country string
}

type Employee struct {
    Name    string
    Address  // Embedding
}

func main() {
    emp := Employee{
        Name: "করিম",
        Address: Address{
            Street:  "মুগদা লেন",
            City:    "ঢাকা",
            Country: "বাংলাদেশ",
        },
    }

    fmt.Println("নাম:", emp.Name)
    fmt.Println("শহর:", emp.City)  // সরাসরি access করা যায়
}
```

---

## 🎯 Interface

### Interface বোঝা

#### ✔️ Interface কী?

Interface হল methods-এর একটি collection যা type ইমপ্লিমেন্ট করতে পারে।

```go
package main

import "fmt"

// Interface definition
type Animal interface {
    Sound() string
    Move() string
}

// Dog struct
type Dog struct {
    Name string
}

func (d Dog) Sound() string {
    return "ঘেউ ঘেউ"
}

func (d Dog) Move() string {
    return "চারপায়ে দৌড়াচ্ছে"
}

// Cat struct
type Cat struct {
    Name string
}

func (c Cat) Sound() string {
    return "মিঁ মিঁ"
}

func (c Cat) Move() string {
    return "সারারাত ঘোরাঘুরি করছে"
}

// Interface ব্যবহার করা
func describeAnimal(a Animal) {
    fmt.Printf("%s: %s | %s\n", 
        fmt.Sprintf("প্রাণী"), 
        a.Sound(), 
        a.Move())
}

func main() {
    dog := Dog{Name: "বোবি"}
    cat := Cat{Name: "মিমি"}

    describeAnimal(dog)
    describeAnimal(cat)
}
```

**Output:**
```
প্রাণী: ঘেউ ঘেউ | চারপায়ে দৌড়াচ্ছে
প্রাণী: মিঁ মিঁ | সারারাত ঘোরাঘুরি করছে
```

#### ✔️ Type Assertion

```go
package main

import "fmt"

func main() {
    var i interface{} = "বাংলা"

    // Type assertion
    str, ok := i.(string)
    if ok {
        fmt.Println("String:", str)
    }

    // Type assertion without check - risky
    num := i.(int)  // PANIC যদি string হয়
}
```

#### ✔️ Type Switch

```go
package main

import "fmt"

func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("%d একটি integer\n", v)
    case string:
        fmt.Printf("%q একটি string\n", v)
    case bool:
        fmt.Printf("%v একটি boolean\n", v)
    default:
        fmt.Printf("অজানা ধরন: %T\n", v)
    }
}

func main() {
    describe(42)
    describe("করিম")
    describe(true)
}
```

---

## 📦 Packages এবং Modules

### Package কী?

```go
// প্যাকেজ ঘোষণা (file এর শুরুতে)
package main

// Import statements
import (
    "fmt"
    "math"
    "strings"
)

func main() {
    fmt.Println(math.Pi)
    fmt.Println(strings.ToUpper("বাংলা"))
}
```

### Custom Packages

```
myproject/
├── main.go
└── math/
    └── calculator.go
```

**calculator.go:**
```go
package math

import "fmt"

// Exported function (Capital letter দিয়ে শুরু)
func Add(a, b int) int {
    return a + b
}

// Unexported function (lowercase)
func privateHelper() {
    fmt.Println("এটি private")
}
```

**main.go:**
```go
package main

import (
    "fmt"
    "./math"  // local package
)

func main() {
    result := math.Add(10, 20)
    fmt.Println("যোগফল:", result)
}
```

### Go Modules

```bash
# Module initialize করা
go mod init github.com/shagorrobidas/myapp

# Dependency যোগ করা
go get github.com/gin-gonic/gin

# Module এ external package use করা
```

---

## ⚠️ Error Handling

### Error কী?

Error হল একটি situation যেখানে কিছু ভুল হয়েছে এবং প্রোগ্রামকে তা handle করতে হবে।

#### ✔️ Error Interface

```go
package main

import (
    "errors"
    "fmt"
)

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("শূন্য দিয়ে ভাগ করা যায় না")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("ত্রুটি:", err)
    } else {
        fmt.Println("ফলাফল:", result)
    }

    result, err = divide(10, 0)
    if err != nil {
        fmt.Println("ত্রুটি:", err)
    }
}
```

#### ✔️ Custom Errors

```go
package main

import (
    "fmt"
    "log"
)

type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("ক্ষেত্র '%s': %s", e.Field, e.Message)
}

func validateEmail(email string) error {
    if email == "" {
        return ValidationError{Field: "ইমেল", Message: "খালি হতে পারে না"}
    }
    return nil
}

func main() {
    err := validateEmail("")
    if err != nil {
        log.Fatal(err)
    }
}
```

#### ✔️ Panic এবং Recover

```go
package main

import (
    "fmt"
)

func riskyOperation() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("পুনরুদ্ধার করা:", r)
        }
    }()

    panic("কিছু ভুল হয়েছে!")
}

func main() {
    riskyOperation()
    fmt.Println("প্রোগ্রাম চলছে")
}
```

---

## ⚙️ Concurrency

### Goroutines

#### ✔️ Goroutine কী?

Goroutine হল একটি lightweight thread যা Go runtime দ্বারা পরিচালিত হয়।

```go
package main

import (
    "fmt"
    "time"
)

func task(name string) {
    for i := 1; i <= 3; i++ {
        fmt.Printf("%s - কাজ %d\n", name, i)
        time.Sleep(1 * time.Second)
    }
}

func main() {
    // Sequential execution
    task("Task-1")
    task("Task-2")

    fmt.Println("\n--- Goroutines সহ ---\n")

    // Concurrent execution
    go task("Goroutine-1")
    go task("Goroutine-2")

    time.Sleep(4 * time.Second)  // Main goroutine অপেক্ষা করে
    fmt.Println("সমাপ্ত")
}
```

### Channels

#### ✔️ Channel কী?

Channel হল goroutines-এর মধ্যে যোগাযোগের মাধ্যম।

```go
package main

import "fmt"

func send(ch chan string) {
    ch <- "নমস্কার"
    ch <- "গুড বাই"
}

func main() {
    ch := make(chan string)

    go send(ch)

    msg1 := <-ch  // প্রথম বার্তা পান
    msg2 := <-ch  // দ্বিতীয় বার্তা পান

    fmt.Println(msg1)
    fmt.Println(msg2)
}
```

**Output:**
```
নমস্কার
গুড বাই
```

#### ✔️ Buffered Channels

```go
package main

import "fmt"

func main() {
    // Buffered channel - 2 মূল্য ধরে রাখতে পারে
    ch := make(chan int, 2)

    ch <- 1
    ch <- 2

    fmt.Println(<-ch)  // 1
    fmt.Println(<-ch)  // 2
}
```

#### ✔️ Select Statement

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "প্রথম"
    }()

    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "দ্বিতীয়"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println("Channel 1:", msg1)
        case msg2 := <-ch2:
            fmt.Println("Channel 2:", msg2)
        }
    }
}
```

---

## 🔐 Advanced Topics

### JSON Processing

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email,omitempty"`
}

func main() {
    // Struct থেকে JSON
    person := Person{
        Name:  "করিম",
        Age:   25,
        Email: "karim@example.com",
    }

    jsonData, _ := json.Marshal(person)
    fmt.Println("JSON:", string(jsonData))

    // JSON থেকে Struct
    jsonStr := `{"name":"সালিম","age":30}`
    var newPerson Person
    json.Unmarshal([]byte(jsonStr), &newPerson)
    fmt.Println("Struct:", newPerson)
}
```

### File Operations

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // File লেখা
    file, _ := os.Create("data.txt")
    defer file.Close()

    file.WriteString("আমরা বাংলাদেশী\n")
    file.WriteString("এটি আমাদের দেশ\n")

    // File পড়া
    readFile, _ := os.Open("data.txt")
    defer readFile.Close()

    scanner := bufio.NewScanner(readFile)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
}
```

---

## 🏆 HTTP এবং REST API

### Simple HTTP Server

```go
package main

import (
    "fmt"
    "net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprintf(w, "<h1>আল-সালাম আলাইকুম! এটি একটি Go সার্ভার</h1>")
}

func main() {
    http.HandleFunc("/", homeHandler)
    
    fmt.Println("সার্ভার চলছে: http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
```

---

## 📊 Project Structure

একটি পেশাদার Go প্রজেক্ট এর গঠন:

```
myapp/
├── cmd/
│   └── main.go
├── pkg/
│   ├── config/
│   │   └── config.go
│   ├── handler/
│   │   └── handler.go
│   ├── model/
│   │   └── model.go
│   └── repository/
│       └── repository.go
├── internal/
│   ├── database/
│   │   └── db.go
│   └── middleware/
│       └── middleware.go
├── migrations/
│   └── 001_initial.sql
├── tests/
│   └── handler_test.go
├── .env
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

---

## 🎓 Interview প্রশ্ন এবং উত্তর

### १. Goroutines এবং Threads-এর পার্থক্য

**Q:** Goroutines এবং OS Threads-এ পার্থক্য কী?

**A:** 
| বৈশিষ্ট্য | Goroutine | OS Thread |
|---------|-----------|----------|
| তৈরির খরচ | কম (micro-seconds) | বেশি (milliseconds) |
| মেমোরি | ২KB | ১-२MB |
| সংখ্যা | হাজার হাজার সম্ভব | শত শত সম্ভব |
| Context Switching | দ্রুত | ধীর |

### २. Interface{} কী ব্যবহার করা হয়?

**Q:** `interface{}` কী এবং কখন ব্যবহার করবেন?

**A:** `interface{}` হল empty interface যা সব ধরনের মূল্য গ্রহণ করতে পারে।

```go
func process(i interface{}) {
    // যেকোনো ধরনের মূল্য গ্রহণ করে
}

process(42)
process("বাংলা")
process(true)
```

### ३. Defer-এর ব্যবহার

**Q:** Defer statement কী এবং এটি কখন execute হয়?

**A:** Defer statement function exit-এর আগে execute হয়।

```go
func main() {
    defer fmt.Println("শেষ")
    fmt.Println("মধ্য")
    fmt.Println("শুরু")
}
// Output: শুরু, মধ্য, শেষ
```

---

## 📚 Study Schedule

### 30 Days Study Plan

| সপ্তাহ | ফোকাস | দৈনিক লক্ষ্য |
|------|------|----------|
| Week 1 | Basics | Installation, Variables, Data Types (2-3 ঘন্টা/দিন) |
| Week 2 | Functions | Control Flow, Loops, Functions (3 ঘন্টা/দিন) |
| Week 3 | Advanced | Collections, Pointers, Structs (3-4 ঘন্টা/দিন) |
| Week 4 | Projects | 2-3 ছোট প্রজেক্ট, Revision (4-5 ঘন্টা/দিন) |

### 60 Days Study Plan

প্রথম 30 দিনের পরে:

| সপ্তাহ | ফোকাস | বিশেষত্ব |
|------|------|--------|
| Week 5-6 | Concurrency | Goroutines, Channels, Select |
| Week 7-8 | Web Development | HTTP, REST API, Middleware |
| Week 9 | Database | PostgreSQL, MySQL, GORM |
| Week 10 | Projects | Intermediate Backend Project |

---

## 🔗 Resources

### অনলাইন ডকুমেন্টেশন
- [Go Official Docs](https://golang.org/doc)
- [Go by Example](https://gobyexample.com)
- [Go Packages](https://pkg.go.dev)

### শীর্ষ Go ফ্রেমওয়ার্ক
- **Gin** - Fast web framework
- **Fiber** - Express-like framework
- **GORM** - ORM library
- **gRPC** - RPC framework

### সুপারিশকৃত বই
- "The Go Programming Language" - Donovan & Kernighan
- "Building Web Applications with Go" - Larry Ullman

---

## 🤝 Contributing

এই রিপোজিটরিতে অবদান স্বাগত! যদি আপনি:
- নতুন উদাহরণ যোগ করতে চান
- ত্রুটি ঠিক করতে চান
- আরও ভাল ব্যাখ্যা প্রদান করতে চান

তাহলে একটি Pull Request খুলুন।

### Contributing Steps:
1. রিপোজিটরি Fork করুন
2. একটি feature branch তৈরি করুন (`git checkout -b feature/NewFeature`)
3. আপনার পরিবর্তন commit করুন (`git commit -m 'নতুন বৈশিষ্ট্য যোগ করেছি'`)
4. Branch-এ push করুন (`git push origin feature/NewFeature`)
5. একটি Pull Request খুলুন

---

## 📄 License

এই প্রজেক্ট **MIT License** এর অধীন লাইসেন্সকৃত।

```
MIT License

Copyright (c) 2024 Shagor Robidas

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
```

---

## 👨‍💼 লেখক

**Shagor Robidas**

- 🌐 [GitHub](https://github.com/shagorrobidas)
- 📧 Email: shagorrobidas@example.com
- 🐦 Twitter: [@ShagorRobidas](https://twitter.com/ShagorRobidas)

---

## 💡 অনুপ্রেরণা এবং উপসংহার

Go প্রোগ্রামিং শেখা একটি অসাধারণ যাত্রা। এই ভাষার সরলতা এবং শক্তি আপনাকে অবিশ্বাস্য জিনিস তৈরি করতে সক্ষম করবে। 

### চূড়ান্ত লক্ষ্য:

✨ **প্রতিদিন শিখুন, প্রতিদিন কোড করুন, প্রতিদিন উন্নতি করুন।**

Go দিয়ে আপনার স্বপ্নের অ্যাপ্লিকেশন তৈরি করুন এবং বিশ্বের সাথে শেয়ার করুন।

---

## 🌟 Stars এবং Support

এই রিপোজিটরি পছন্দ হলে ⭐ star দিন এবং আপনার বন্ধুদের শেয়ার করুন!

**Happy Coding! 🚀 Go-এ দক্ষতা অর্জন করুন এবং একজন পেশাদার Backend Developer হয়ে উঠুন।**

---

**আপডেট: July 2024** | **Status: ✅ Active Development**
