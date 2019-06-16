
  ``` Go
  
      package  "text/template"
  
  ```
  <br/><br/>
  ## Template
  [template.Template](https://godoc.org/text/template#Template)
  ``` Go
    type Template struct {
        *parse.Tree
        // contains filtered or unexported fields
     }
  
  ```
  - pointer to template is used to store multiple containers and that templates get parsed once and used later whenever required
  - templates First gets parsed then Executed <br/>
  - To use templates we first need to parse template using function `ParseFiles() ` or `ParseGlob()` 
  - Both `ParseFiles() ` and `ParseGlob()` returns pointer to template i.e. `*template.Template` and error.
  - The difference between `ParseFiles() ` and `ParseGlob()` is `ParseFiles() ` takes in argument files name
    while `ParseGlob()` takes pattern of file. e.g. <br/>
   ``` Go
         tpl , err := template.ParseFiles("one.xyz","two.zyx",..,"n.abc")  // extension of files doesn't really matter here 
         // or 
         tpl , err := template.ParseGlob("*html")
   ```
   - Once we have pointer to template then we can use following Methods
   
   ``` Go
             func Must(t *Template, err error) *Template
   
   ```
   It is used for parse files in a better way. i.e. you will get pointer to template on sucessful parsing of files and any error in parsing will handled by Must() function.   
   
   ``` Go
             func ParseFiles(filenames ...string) (*Template, error)
   
   ```
 
   
   ``` Go 
             func ParseGlob(pattern string) (*Template, error)
             
   ```
   ``` Go 
   
             func (t *Template) Execute(wr io.Writer, data interface{}) error
         
   ```
   If your pointer to template have multiple templates then Execute() function will execute first template. 
   ``` Go
             func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
   
   ```
   It is used to execute one template you have to specify template name.
   
   ``` Go
             func (t *Template) Funcs(funcMap FuncMap) *Template
   
   ```
  It  is used for pass functions in template , you can also use global functions that are already provided. 
  
