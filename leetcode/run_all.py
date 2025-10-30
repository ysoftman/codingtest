import os
import subprocess
from multiprocessing import Pool


def go_run(dirpath):
    try:
        # go.mod 가 있는 경우만 go run . 로 실행할 수 있다.
        # subprocess.run(["go", "run", "."], cwd=dirpath, check=True)
        for _, _, filenames in os.walk(dirpath):
            go_files = [f for f in filenames if f.endswith(".go")]
            for f in go_files:
                print(f"▶ 실행 중: {dirpath}/{f}")
                subprocess.run(
                    ["go", "run", f], cwd=dirpath, stdout=subprocess.DEVNULL, check=True
                )
    except subprocess.CalledProcessError as e:
        print(f"❌ 실패: {dirpath}")
        print(e)


if __name__ == "__main__":
    root_dir = "."
    dirs = [
        dirpath
        for dirpath, _, filenames in os.walk(root_dir)
        if any(f.endswith(".go") for f in filenames)
    ]

    with Pool(processes=20) as pool:  # 동시에 20개 실행
        pool.map(go_run, dirs)
