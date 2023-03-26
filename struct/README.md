## Notes: 

### Example App Scructure

```
human_main_dir
├── go.mod
├── main.go
└── person_module_dir
    └── person_module.go
```


### Pattern

* A module is referenced by a folder name of the module
	- Example: `example.com/human_main_dir/person_module_dir`

* File name of the module i.e `person_module.go` should not be referenced anywhere in the app. Instead, `main.go` calls the module using `person_module_dir`

* Module dir name and module name witin the module file i.e `person_module.go` should match
	* This is a tricky one. I tripped on this twice 
	* Always name a logical name within the module file as same as the module directory

* Witin go.mod 
	- Anything before before the `/` doesn't matter in `example.com/human_main_dir` in `go.mod` when the app is not importing any local module 
	- It matters when you import module 
	- Import path in `main.go` should be `example.com/human_main_dir/person_module_dir`
	
* When you import module, variable names should be all in upper case



### Go mod 

`go mod init example.com/human_main_dir`