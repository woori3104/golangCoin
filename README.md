# golangCoin

### Variables in GO

- memory 위에 관리되는 값을 저장하기 위한 영역 
    - 식별자 (변수명)으로 식별됨 
    - 처리중의 일시적인 값을 저장하기 위해 필요 
    - 메모리위에 있는 값을 머신이 멈추면 삭제됨 
    - 프로그램의 실행이 종료되면 삭제됨
    - 결과를 남기기 위한 경우 파일등 외부기억장치로 값을 저장할 필요가 있음. 
    
- Go의 경우 변수의 타입을 확실하게 지정해줘야함 
    - 어떤 종류의 값인지 지정해줄 필요성이 있음 
    - 유저 정의 타입을 선언하는것도 가능 
    - 처음 정한 타입이외의 값을 대입하는것은 불가능 

- 정적형식의 이점과 Go의 타입
    - 컴파일 시점에 타입이 일치하지않는기 검출 가능 
    - 암묵적인 형변환이 없음. ``` 1+"1" = "12" ```의 형태는 불가능함. (자바스크립트 가능)
        - 암묵적인 형변환을 하지않음으로 유저가 의도하지 않은 동작을 사전 검출 
    - 타입 추론이 가능 
        - 타입을 무조건 지정하지않아도 입력된 값을 기반으로 추론하는 것이 가능함
        ```
        name := "kim" // string
        age := 20 // int
        age = "12" // compile error : age에서 int로 추론된 이후 타입을 변경할 수 없음
        ```
        - ``` := ``` 과  ```=```의 차이 
            - ```:=```의 경우 변수를 만들고 대입 ```=```의 경우 값만 대입
        - 함수 밖에서 변수를 선언하기 위해서는 
            ``` var name string = "kim" ``` 으로 타입을 생갹할 수 없음
- 정수 
    - 바뀌지 않는 값  ``` const ``` 
    - 정수의 이용
    ``` 
        func main() {
            const (
                a = 1 + 2
                b
                c
            )
            fmt.Println(a, b, c) //  3, 3, 3 
        }
    ``` 

### Functions 
- 인수를 받아서 처리를 실행해서 반환값으로 결과를 반환하는 기능 
    - 반드시 인수나, 반환값이 필요하진않음 
- 인수 : 함수의 입력, (파라미터)
- 반환값 : 함수의 출력값 
```
// case 1
func plus (x int, y int) int {
    return x+y
}
// case 2
func plus (x, y int) int {
    return x+y
}

// 복수개의 리턴값을 갖는 경우 
// case 1
func plus (x, y int, name string) (int, string) {
    return x+y, name
}

// 인수의 개수가 정해지지 않은 경우 사용할 수 있는 방법
func plu]s(a...int) int {
    var total int
    for _, iten := range a {
        total += item
    }
    return total
}
```
