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

### Package
- What is Package 
    - 함수, 변수, 정수, 타입등을 의미있는 단위들을 모아둔 것
    - 한 파일로 작성된 패키지인 컴포넌트는 코드의 모듈화, 코드의 재사용을 가능하게 함 
- GO에서의 Private, Public
    - 변수의 이름이 대문자로 시작하는 경우 Public
    - 변수의 이름이 소문자로 시작하는 경우 Private
- 파일스코프 
    - 파일내에서 임포트한 패키지를 저장하는 스코프 
    - 패키지이외는 대상이 아님 
- 패키지스코프 
    - 패키지 단위 
    - 대문자로 시작하는 식별자는 다른 패키지에서 참조 가능
- 패키지변수 
    - 함수간에 변수를 공유하는 방법 
        - 함수내에 정의된 변수는 함수 밖에서 사용 불가 
        - 함수 밖에서 변수 선언 
        - 함수 밖에서 변수를 선언하기 위해서는 
            ``` var name string = "kim" ``` 으로 타입을 생략할 수 없음

### Slices and Arrays
```
foods := [3]string{"1"."2"."3"}
```

### 블럭체인의 연결
    B1
        b1Hash = (data+"")
    B2
        b2Hash = (data+b1Hash)
    B3
        b3Hash = (data+b2Hash)

    - 입력값이 변경되면 출력값도 변경 
    - Hash의 중간값이 변경되면 (B2) 그 이후의 블럭은 무효화
    
### sha256
- 미국 국가안보국(NSA)가 설계한 암호화 해시 함수들의 집합이다.
- 해쉬 함수로 임의의 정보를 고정된 길이의 암호문(메시지 다이제스트 등으로 불리기도 함)으로 반환한다.
- SALT 값으로 해쉬값을 만들고 DB에 저장하면 관리자도 비밀번호를 알 수 없게 상태가 된다 
    - 복호화가 불가능한 알고리즘이다. 
