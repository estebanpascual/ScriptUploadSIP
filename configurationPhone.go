package main 
  
import ( 
    "fmt"
	"github.com/imroc/req"
	"github.com/harry1453/go-common-file-dialog/cfd"
	"github.com/harry1453/go-common-file-dialog/cfdutil"
	"io/ioutil"
	"strings"
) 
  
  
  func LaCombo() (combolist string){

	result, err := cfdutil.ShowOpenFileDialog(cfd.DialogConfig{
		Title: "Ouvrir la config",
		Role:  "Ouvrir",
		FileFilters: []cfd.FileFilter{
			{
				DisplayName: "Text Files (*.txt)",
				Pattern:     "*.txt",
			},
		},
		SelectedFileFilterIndex: 2,
		FileName:                "config.txt",
		DefaultExtension:        "txt",
	})
	if err != nil {
		// log.Fatal(err)
	}
	
	return result
}

// Calling main 
func main() { 
  
    var ipAddres string 

	fmt.Print("IP du téléphone : ")
    fmt.Scanf("%s", &ipAddres) 
	
	nomConf := LaCombo()
	content, err := ioutil.ReadFile(nomConf)
	lines := strings.Split(string(content), "\r\n")
	
	Nom := strings.Split(lines[0], ":")
	fmt.Print("Nom  : ")
	fmt.Println(Nom[1])

	IpSip := strings.Split(lines[1], ":")
	fmt.Print("IP serveur : ")
	fmt.Println(IpSip[1])
	
	NomSip := strings.Split(lines[2], ":")
	fmt.Print("Nom SIP : ")
	fmt.Println(NomSip[1])	
	
	AuthSip := strings.Split(lines[3], ":")
	fmt.Print("Numero SIP : ")
	fmt.Println(AuthSip[1])
	
	MdpSip := strings.Split(lines[4], ":")
	fmt.Print("mot de passe SIP : ")
	fmt.Println(MdpSip[1])
	
	NomAffiche := strings.Split(lines[5], ":")
	fmt.Print("Nom Affiché : ")
	fmt.Println(NomAffiche[1])
	
	
	req1, err := req.Get("http://"+ipAddres+"/manager?action=login&Username=admin&Secret=admin&time=1608192935197")	
	if err != nil {
		fmt.Println("Found an error", err)
	}
	AllCookie := (req1.Response().Cookies())
  
	req.Get("http://"+ipAddres+"/manager?action=put&flag=1&var-0000=271&val-0000=1&var-0001=270&val-0001="+Nom[1]+"&var-0002=47&val-0002="+IpSip[1]+"&var-0003=35&val-0003="+NomSip[1]+"&var-0004=36&val-0004="+AuthSip[1]+"&var-0005=34&val-0005="+MdpSip[1]+"&var-0006=33&val-0006=*26&var-0007=3&val-0007="+NomAffiche[1]+"&var-0008=63&val-0008=0&time=1608192935197", AllCookie)
	req.Get("http://"+ipAddres+"/manager?action=needapply&time=1608192935197", AllCookie)
	req.Get("http://"+ipAddres+"/manager?action=applypvalue&time=1608194015451", AllCookie)

	req5, _ := req.Get("http://"+ipAddres+"/manager?action=applypvaluersps&time=1608192935197", AllCookie)
	body5, _ := req5.ToString()

    fmt.Println(string(body5))
	fmt.Println("-------------Fin de l'upload-------------")
	
} 