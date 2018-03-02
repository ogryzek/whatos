/*
Package whatos provides basic functionality into returning what OS your are running, provided it is one of:

  * macOS (darwin)
  * linux (linux)
  * other (anything else)


Example:

                import "github.com/ogryzek/whatos"

                func main() {
                  myOs := whatos.TellMe()
                  fmt.Println("My current OS is: ", myOs)
                }

*/
package whatos
