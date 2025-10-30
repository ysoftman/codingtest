#!/bin/bash

# golang 코드 실행시 에러 발생하는 코드 찾기
for d in *; do
    if [ -d "$d" ]; then
        pushd "$d" >/dev/null || (
            echo "error pushd $d"
            exit 1
        )
        for f in *; do
            if [[ $f == *".go" ]]; then
                echo "실행중: $PWD/$f"
                # 많은 파일의 빠른 확인을 위해 백그라우드로 실행
                go run $f >/dev/null || echo "run error $f" &
            fi
        done
        popd >/dev/null || (
            echo "error popd $d"
            exit 1
        )
    else
        echo "$d is not directory. skip"
    fi
done
