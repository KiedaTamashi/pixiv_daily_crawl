package setting

import "fmt"

const (
	username = ""
	password = ""
	cookie   = "first_visit_datetime_pc=2021-12-31+14:04:47; PHPSESSID=dh3sov103o6h73djd2igqpija4vorct1; yuid_b=F4R3R5I; p_ab_id=2; p_ab_id_2=8; p_ab_d_id=992184231; __utma=235335808.2108051685.1640927088.1640927088.1640927088.1; __utmc=235335808; __utmz=235335808.1640927088.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); __utmv=235335808.|3=plan=normal=1^11=lang=zh=1; __cf_bm=1thgJlp9keVIqEjghpI4SpzbE3wA.RHAcD2RHdZW35I-1640927088-0-AZdhRsrmDlUUa5MYAUjwCrFaK/SjOrUWdPWE2lXguLI6IahXPyqWz2lLjb3/bBWoZrzNs5tISVpJ4W0ReWryO+u7woqATHvkOv7E/86W5G5/gTLKgIyig+PhICxHfy9WQaSnA5lwuu3emFYGqorV0Sc1Wl0MRpFgH+UnBtxmf6OT+uzhhPr2dlWJuDKzIuXQTA==; _fbp=fb.1.1640927090388.1859668556; tag_view_ranking=yPNaP3JSNF~jyzj4lA07D~4ZEPYJhfGu~hF2c0817Ah~4VJIHse340~RDw9eirzSy; __utmt=1; __utmb=235335808.1.10.1640927088; tags_sended=1; categorized_tags=yPNaP3JSNF"
)

func GetUserName() string {
	return username
}

func GetPassword() string {
	return password
}

func GetCookie() string {
	var cookieDynamic string
	fmt.Println("从pixiv输入cookie（不带引号）：")
	fmt.Scanln(&cookieDynamic)
	return cookieDynamic
}

const (
	DataFolder = "./data/"
	CsvName    = "./mapping2021.csv"
)
