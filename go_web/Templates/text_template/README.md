
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

  - templates First gets parsed then Executed <br/>
  
  - To use templates we first need to parse template using function `ParseFiles() ` or `ParseGlob()` 
  - Both `ParseFiles() ` and `ParseGlob()` returns pointer to template i.e. `*template.Template` and erreor.
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
   
   ``` Go
             func ParseFiles(filenames ...string) (*Template, error)
   
   ```
   ``` Go 
             func ParseGlob(pattern string) (*Template, error)
             
   ```
   ``` Go 
   
             func (t *Template) Execute(wr io.Writer, data interface{}) error
         
   ```
   ``` Go
             func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
   
   ```
   
   ``` Go
             func (t *Template) Funcs(funcMap FuncMap) *Template
   
   ```
   ``` Go 
             func (t *Template) ParseFiles(filenames ...string) (*Template, error)
             
   ```
   ``` Go 
   
             func (t *Template) ParseGlob(pattern string) (*Template, error)
         
   ```
