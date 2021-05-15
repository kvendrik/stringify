run_demo:
	cd demo && tinygo run demo.go

check_wasm_size:
	tinygo build -o stringify.wasm -target wasm ./stringify.go && stat -f%z ./stringify.wasm && rm ./stringify.wasm
