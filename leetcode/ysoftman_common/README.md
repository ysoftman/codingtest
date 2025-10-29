# ysoftmancommon

데이터구조나 디버깅등의 기능을 모아놓은 패키지

- 사용하려면 프로젝트에서 다음과 같이 추가해서 사용

```bash
# go.mod 에 다음 추가
replace github.com/ysoftman/ysoftmancommon => ../ysoftman_common
require github.com/ysoftman/ysoftmancommon v0.0.0-00010101000000-000000000000 // indirect
# require 는go get 으로 추가해도 됨
# go get "github.com/ysoftman/ysoftmancommon"

# .go import 추가
import "github.com/ysoftman/ysoftmancommon"
```
