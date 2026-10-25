package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "3.0" )

func bbQP0HCqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nr0urETZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J2SkqTleWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dHfcXuxyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xwkBAkKAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YLHDu4xxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FTnE2cqxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Ts8X7kGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QOFJgnEFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ANcWem1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uySlywPZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SFrk6HxaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func an7ZVYwDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BV5c13MQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MDUrUYWGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OJtP4DniWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WPprv61XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QJNtvsp9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jdNOXyD3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0I3vt15tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b8BYSp60Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CtcwFxJwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3UNrZCKwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3jNPaVVcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MlCziYfzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iAlTn02jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func azUaLjkaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ME0c05kkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W9kTWjJXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6B7u5K1qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kAHsfWMVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vfc56NFlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 82rgP6xWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C6AGMTRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func km9lP3MAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kbd3SN1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func czVTCOnbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kLCpaRQ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o4wbBsMGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FG8hhvJwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6DqMVeAAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AF8AD0WgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mX2Jm0HfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FqG7STGxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uM2uZsZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rVBQloqFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func upMjzOS7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PNhNKBvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x0YeLCKHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JVonHlNFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zGH92NQ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func caKyyxMDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 36y3s7WrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tp2zOsdAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cNxQdfd7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nghJSjWeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CglrjISSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YinGATlKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aGSNRlCnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uk5hzpxOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vyHWozdgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b0OFoYxIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iUgA4Wr9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HhoWh7oiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gWiWHNaPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eloFGanWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IH4iwri4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SuV58W8NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dicZfr3jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IufSaaEfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PqLcDFLuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q1QjxvrUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KglRFsziWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ObUCS53eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n3FCR4b0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ii3QPtydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BC8Uy7BrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H1pLsZgNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8qf72wDXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CamjYQIcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dA7jTMmCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eLQmnry9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HgHIQvDkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oenkWDS6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kat5zJcZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JR6pDRNfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JTNgRDdIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GDGKL4nTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ohdEnFJ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P0otZUWJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gfjs9vNYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lqYzXTC0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qs7O8YmyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zDHdd6RgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6BYcpU4oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sn0tAccbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 03d5jZWRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R2JLdSRFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6dBoTA9tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4gUaSaw2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ja2XdS1JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UozGu5cqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jeu3eOj8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CPOOImISWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ow2lgFzyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F8zDQbGmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W6UR3xgTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DXdCJZjjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MvtzllxuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q6mLeDURWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a60QSJWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tbU5reDcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iBMfzmmbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u8pEqow9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8tZNX6RpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XStgw9zjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fxOHRu8MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func znAC3gKNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WesfgKJnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IEaewQ1yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NohxydUuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ppMM4oLwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7sUIFYZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FrhPyOrAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EobbkpFcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZWRsCpyfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func toPQ6zCzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gGbulvIqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0v9MBVOWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UtVSr39DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5DmQb7jWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yzDmE4xLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NyIMu1YrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e3q3gmcOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1arFi1YuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2S7vHzRzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vMb2JFwGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tHvHx7tMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PKfKHlZtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vd0HzOs1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hNkSij49Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y85AQbaxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eWnhpL0BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RPwMh1SrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y9df9YldWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3qgyo59bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vdpiaLkOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O6QMTod5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LWbzgSfsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K7upJlNXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TtMzmVd3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KUOwE1CtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uPKNjDJdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VlWLwjEdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mB1m6ogGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kk6Is3KnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IdS9nZFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P6l2kksVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KRBBGqmMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L2QBziRcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hSsQUVp3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S56rlwW6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Kp3dSvHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n0jFvYaYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PToGVAc7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kmeX1uSaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PwKTCQ1PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H6vb2yCEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uZfjN8DLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kes5aEerWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iuikcL6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OZYxfOvEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W4uSje2OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IQIYaJ4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ke62wlfSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wXmzwx0UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NxIOvCPZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CoziHoHMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9hatetGeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nIu6CpHJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dGaKBonoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E7vhuiKNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vrvQhe9CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NUjBzLNAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vN9i7mN2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YjqsBwqUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4MXH0bLaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EZbiRFqmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SYD4Z8vGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PzzrKfg7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Ty67cnPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bl6GkRT7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NTX8hhPTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cFnYCYSxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jLPn0UESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZX0uTAMyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B8JOOzOuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OhcgexfwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iuVrhFMpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dsAu5XjqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gehlr1fVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1y1rt4NCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2CgPH3MVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fg9JYxRJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9CpFyYKOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ywzlnq84Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C4O3mldQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hwgaeWkwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0jaitfDMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5YhE3kwDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UsDd3QqfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d5Z1Sfq5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pKWXlKkvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DcjHk6DSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TfFkpy6tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zVvInFHPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f3XARCFOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4KxFd7uoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IwDQvLHAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8CZzhGAjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vPFapmPdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6X1FGHQxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p4YCCnxaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DbLwtmqmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hrLulAwuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s4QGBz8ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DvJzVHTcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hBmc3xkvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MVyxZfhWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lK5AWsMbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mxrssv1wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9V446suOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uI3ctZboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IJoaBI9LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xlbU2z3jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dXaivjyQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a1vncHJAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4gcHgu9RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P7DUfNJKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TIMfH4VAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mcp6X4JOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rZXYPB8QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HRgbutLqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pjXs2gIGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f9EBm4QFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jDlusQ2WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oSEWzeqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yHVD1lP1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K5VqbTalWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4vGJt27bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IVcibuw7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gVJIPPhzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KsGqYmnEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JRRUcrzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GsUyvj1LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gi71NgmSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Suf764uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z4IzFZotWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZgQRYKI6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5bWhrZlwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jXfWLt9qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IOZxs0X9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l7ODn9mBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sMnVru8BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A50NYb10Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mxkSmL5BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XjgaEWl9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e6DwUuG7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kveKFCL3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EWLAxJVaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BpdCj1UXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HUxFmdCxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8r6iqWGbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 348jyZuFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vl7CTduyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M3r8FhBSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9dRxVPqbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pi0ZiUJ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3mkMIaPgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tY1o12iRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ti4lkD1AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wq8Xvnd6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e8nM09lXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rrsb67MeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RcFAyHKpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TIE7BXyFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T3eEGI5BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VsPYeWo5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FsMASz8vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func idYxvUY6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uBiO4lUDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9C0VXQ09Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TTPxegG5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xTZKNwIWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ReeNuWDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AxuI0PGZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QVToUAi7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FwzF1yCBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6o3Lu2h5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l4CWrsPZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p8bVsqahWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func smr3Zm9fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J1lJhBRsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wdCaFpPrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ockM7LvtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vOYRHo97Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bEYVWYRXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fne1anSFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DjjwVIiAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IDwFqtQWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ODen8M8kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BCS75ZcgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jgv2vtvJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UzZe0UMeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d4VVnTBoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iya0zX3qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EgPodUMhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FDHNVZBYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oQP5r8w1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hxrCn7dMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JODHV8SjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MBlakzEiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nlXiSPsEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cow9jkkwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j980YqBuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GXgMuefgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3DGdQSWYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bQ0z5ucbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wBqXFuzbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7WjWmpFdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jDEd69AWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4PAqcJsdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QZrNbe98Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I87RvIjiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KdEEyoe2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5yY9Oa7tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h52flkNjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZL4JQYgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q79qXL8KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p35x3hrRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zIf1lRcNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eyr0I0jQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ylT9PZuOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iOqfIri4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lad8KQY5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HcSnnwU1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8mKycpJYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hIT7ZpOmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K94lKKULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JlhyoHPIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tIPKm3D2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u7D6fY2rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kYKCVrssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Z8mLuQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CF1oAps6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z4o4JLcpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iOubWIomWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OwNkOz9kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qreD0PPoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PMtuXHs0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eibSMU2GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jjLttzpcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9bdVXKdcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MeRvIFZzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IbhsF2jBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1li6zFrmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u2cVdAC9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gG2gF10IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nYmX63RTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vJ8QfmzEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qylxa0zSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xkByOJOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RTUme67vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D0JSSDDUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cgPLxQBIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L1yA6niqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a5RBaYgDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ddGXjrJCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mgP68TGlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CWJbqogkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IqTIkCFPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V9Ll8eOyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VoZdsFVuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uiedeUDLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rPTcT6QpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yVMPkcwhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C7JvvsONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lWlXnm9YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1QkZAktPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uOTk2lnAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ziDmpa7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fzLYvSJ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nWoegnUZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func McQ0cFkeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KZQDFBtAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LW794nIBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9jDyvLjYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9nwOCWpLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kJps6VtEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SvarCvqzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UCZItpyRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C094XXwAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AlqH5M22Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yEBOmt7VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RrPxYGivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tq8rSAfxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V8x4C2NQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ylmychsFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SekGGl5UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BFZ7DLToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MjkLkqNLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s6uc0HWCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hctll3MYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RUZO8GNvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Og8fEO81Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ksnhN5evWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KsXz3rS6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func swPz3YDSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L60r6peLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EmGY6t5JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ctJugOvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fDxBNDnBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oLAxlWbwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ezQnp1ctWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GQbVAWclWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fSaSyzelWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PGALqTUnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6reisWoIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nG2cQcCiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N3zCvk6aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hPSFHzFDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s36PZH3jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XzmqKGpBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zNbdbWpPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bkf3vAZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QXFAI7sUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u8rWNTKdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XSYC7Yr4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JMM5IZoJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p7gzMhGeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pkbQKIqBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jDfu7AnhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I6OabCmbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kMxrtUbIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XmyHxqLxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wcyid45AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ChH8iiajWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dtz6jVYwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TFHEUsfOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func afld9AabWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jpj9yXKSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4c8WID2fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NGhR2Ij3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0FI9MVWXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HJGpYaOcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cmAIgYvcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6wwLdMEmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s2wpkQEEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W2kPLDunWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AUWFoOVCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XpR8CDBnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oty7GK3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GGNsghQNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2hBqIv61Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X0vqSh60Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NlPPYYgnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jrQpBVZnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XQyDeMPvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fovwefT8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HLGQkAnGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I3mfDV7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PzYfhykNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4yHHaO17Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TUirmrF9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XxSM1kULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e7CbXlI9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GJUT2VbyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ARtDVJt6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9EajxFqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nf9cI5TkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fzPGsHiEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 79kVhQbtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I9lpxhDPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6fOLTMckWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zjtjYrUcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sqPfI3W9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y3ieiFhzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N1YxyB1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yDqGJbw0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wLtJYSJWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6TqArT7kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NnEZb9WZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

