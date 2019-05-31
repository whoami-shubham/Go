
 # ` package  "text/template"`
  <br/><br/>

  - In order to serve a template  we use this package
  First we parse the template files then execute the templates <br/>
  
  - To use templates we first need to parse template using function `ParseFiles() ` or `ParseGlob()` 
  - Both `ParseFiles() ` and `ParseGlob()` returns pointer to template i.e. `*template.Template` and erreor.
  - The difference between `ParseFiles() ` and `ParseGlob()` is `ParseFiles() ` takes in argument files name
    while `ParseGlob()` takes pattern of file. e.g. <br/>
         ` tpl , err := template.ParseFiles("one.xyz","two.zyx",..,"n.abc")  extension of files doesn't really matter here 
         `<br/>
         similarly <br/>
         `  tpl , err := template.ParseGlob("*html")
         `
