// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package include

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("journalbeat", "fields.yml", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJzsXE1z5LbRvutXoPRe3lSNKGkka9c6pKLsJraSbLwV21W5jTFAkwMvCdAAqNH416fwQRIkwSFHMytvUtbBNj/Q/XSj8XSjifEF+gS7e0REUQh+hpBmOod7dP7O3kB/E5XkOF8D1udnCFFQRLJSM8Hv0R/PEELoneAaM668CJQyyKlC+AmzHK9zQIwjnOcInoBrpHclqOQM+dfurYgLxHEB90gCpivNClAaF6V9FFXp/n7YADIvo+0GOkCtHKQ3gH52NxFwLXdJRxcREmjVaDG47lEmRTWp96/OwEoBResdUjuloaAXtUD0CSSHHG0gL0EmfmBobgij4kwHsh2MT7DbCkmD+6NgEHrgXGhsnigkUlSAUjgDMx12XhjPGlsVSqUoPGKrWiVDRArk6reAZRSHoGpAfhLpi2eqlOKJUTdbtbCpeRHrn4EMXRCqnXBAo14LlIsMCY7WsMF5atyBEWVpCtIsiVKKTOIiCUb3UYXIcEU7czMObhKgnScjrl60InXLNOm9GMMTYspFxviqYjSiwGHLBc8iDyX8UjEJ9B6lOFcQeQOecVEaRrq+urqKPJ800BGFRYh+fHxvbDTU4CbYOJ+AUn2DW9MUKMUsM34uw25ebpWNhRriHMsa9iv6UzW2xKeNaQw5v8zZ+tLzYf1vdHFhFvb5YXFpjDP5BHOKcsahtm2vUeafr2TVgeb8ExeNCfAMpNImM8ataJ+fzJYJdB+x3hiWMuj2ITzMKU0+GbDC6MI5yoqZa7uGlb0SrG8Og1W+EqyPh8HyM3y6tPO9Z4cjE4/YclOtfI7EM5OCw7AzYF6eSeIr+6RgW6w1O3tQ46h7leCrQ+YBddagDaY9iCMF7JcBuy5wh0W32zN8ieUmhSdGXi0VGdf5/ZNTbFGMcFK1dp79DcA1uvfgcwas1K7IGf+koig1PPfD9CiID5Qy8wznyOu1YBQqBePabLp8nq+9K6jdoF9SeNprhXlxVWK9eRUzfughNIrdamKqfjAL9VhF+DkAD+MWYYVMfb4RW4Wq0kAOfK8lAFpDLrbIVFBDUiCCwkkowfOQkYcy4CCxCwVT6/Wz7j4uSNlIUXqwO9tCMkkulSSXREi4LDDHGciEvGC30CFcUUkCFi7abkC6stawIlO1A4DGIyetONHDNH20nT+L9SoX2UpprCu18v2QIw2twaLthpFNa1pjr1cTN9VsrE5TbLab9GV/NzvDIrvB41WxBtkJ1dlGNcWzK2G/xDT6xXZt/quaNn5h/w82bfZZ9nvT5jRW/d60md20OVkbomWCw2OpXfc+gnxGyEWW+XSwN7+drMVzGiMs0XsTZMXV9No4WevsNAbY/esh+Aku8ZrlTEN/93N0sQ1pCkSzJwiVRAkn3sI6/CuSSGsRU9+OGH8SxH76WnUmcIwLxp3RLt23t19dpdd3b5YU7m7vyNu3hF7f3GBM6W26pG+uzueVN8Z1LTwzt6mQ1mey4vZjKtmRvKG9ijMdrjO0xUEBjRiPbFr69cwxVl+akEtUzgjY/7y4Xt7c+mufQC+WiSKihAMcQATXUuR+Qdq9pd+Z1elvw0BiSTa7oX2xjmN0VY7bNwHPauhUPf3+ERJyvIE3XgMdPhMTSGe0Exs0ebeLdExU9CLhgJlvYJpxI6042z48Hu5MJHZO98KZ9yV+jt94xvhzokA+Hea16RbsS84OfN6ZPrD/qiXmqhTyMOBaVnHcaqdykc2E+63Y9re2lmclEGBPsWMLNeqNUPqYfGa0GhlTCW0thI6lsl47ZM6MUvqWfP3mFiuaXl3TNSwhXd7RN6m5sby7JV8fMMcGVpjC7HXtyXimCoqBXGTH+m6yoTbm0FIyIZnevTyHRCu6CX/VWmv4dfGMHrw/zO6txJrZcmrn0A/XSoqJff664GutR4JnFLhmKQN5imhW1SF1V6P6BTbU+P2QsxHgUW4aweU7LGYr1dlGNZVcpbQoOpo4KA10oCt00IiyB7lmWmK5q5tWRBRmydiTX75eB5WcnfnziGvAuj2N+Gd3NeP0oT3/d+ARxI7VRkASNCX2OC/MMQYhUsBpzQMBm6sEPQZv2WGs3REp0HX9SQRPWVZJV52nLIeFuW8eYo2ecF6ZkfbooZXJtLnkQofCFg2ve03+/R+EVdXBsTDP7K2fzOVPjRxhLR7HlQydVmucdlyDzWY5XUnuspz9iF7WbOo/sDXd7QZ44DtZcc54FkFjNjO/Cj4DTf3m50TzBDIoxveA8S/WYWXD2U5+2AZnqk4zHU3nf2rOzp53li3FGvYRRCpkgXXnvYbjHqqsUhot7/QGLa+u7xboenl/89X9VzfJzc1ynnctJHdMt0mRdoFIIELSbq7uGaVxpuZSi3nXeYtgQwU23kuQbqIwp/bC1nq4+5HH+Kmn2LFDx4+dQ6HuYhWrEkeANlxlC9FmTRmCcsp6CEBKIQ+tUf5iBrUk2xy2xe2XYcZTYVY2wcryl9WjpmqWbuZBY2lzTxJ00Ea/3kQ/d/ZqiknpRkikvN+Vc3Zzk9J9mPgchXOGVZukHvxlRIp9VE9KamGGeX7L9Aa9SZ775+L/D/lT+CZ+VRiR4Q6gT7zOuhqc+Sux3sRIugk0SyY9xtwrpnm3xpmLitZ5vIm/bjlhXkn8WWw5oaQAjZPIiK4wxpXGnEDC6Hx59aB6QzMicoZDY0IHvnXPC0w2jEMSBOIMqX7UqhnVFerrGxtDqxkzF0iODx1MlWG4g5zrx8R9KyFrOXeGMP++D7H3gnwCORFjju9AToOmVlwyGDEUNSMSBsKGYdDqKQz7vUSoHVmTj3VRSz7NArRrxXqRYo3jdPTBP3XVN+kMVSZTtAUQpnRlX1jVItsZGK2hx1ZvUFYAmagdwq9vXYQJ+iiUYiZt2opYISzBCFygjMACCYkoy5jGuSCAeTKKrc8Eo1ge/YtBw9PwKKpX9bSG6bq40RHuKuZpGdBE4Ge9TAqgrCr2a//gRNhYPEz5GAk1CCp1AVjpi2syUcYFgpCtx1lbazPl4DDVFtl7Qi7koACKf3LxPD/0/BCD5RshshzcShvX3iG5EQX/su9M2ecXuqOBdqW/r68jwj1HKm3KBSLyHIjZMdhl7p6ZNas2QuqVqz/bzTvmZCNkre+iWeUjv9VrYKFodTpWRUYYGr2sIvuRs18qaAUiFmn39MnzKI1hXFhx9d7YAzDbmHXFco1iX0CiCeWFSN41OrvHT4e6cryGfHgiavDztj27mQksj9YTTk8TtL4z7UP2W3cVEfJotiJBoPpObpd62tg09ycjc29XfCwuj5+Tb31lHes5niTSHUFEghxLsmEaiK7kCWzoiEP/D0mWoOe3d6u72wXCsligsiQLVLBS/WFP4LFfYWj1WogcMJ8ZV0EWYgrhIKgHaoVKyhzrVMjiOAd89z2qBXnTCXAt1AJV64rraoG2jFOxjdku1AlW9nf9hpMtBxyU8w+YGIT/Po9rN9xzZKh9971lMGp7BLJw/T6v/frt++urr0d0dxtcL9fu5UR1pLhg+fCzw6EqnBhvlQS6wXqBKKwZ5guUSoC1ovvmmQ1PN7CZp0v/wZQ2GeTx4wWmVPqed19BgclxRtZqNljSLZbQKlugSlU4z3fow8O7EIPn7U/VGiQHd2DHs/ffw3sRre3zZtvR3UO0QlHI3fvLkHbQJOF3QKODaL8U9ASLNvBAKajLJVFVsWMjL9X0UVD04+P7oSL7g48Szzq/ME9VK3GoTFA4rQftDz7iLpxbzMxT5KShApdDTbj9vxecSl0gMq7zlAVioJd0asV9ak9QIkf1Orn/CQAA///R3e1k"
}
