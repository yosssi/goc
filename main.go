package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	expectedArgsLen = 1
	baseUrl         = "http://golang.org/pkg/"
)

var (
	pkgs = []string{
		"archive",
		"archive/tar",
		"archive/zip",
		"bufio",
		"builtin",
		"bytes",
		"compress",
		"compress/bzip2",
		"compress/flate",
		"compress/gzip",
		"compress/lzw",
		"compress/zlib",
		"container",
		"container/heap",
		"container/list",
		"container/ring",
		"crypto",
		"crypto/aes",
		"crypto/cipher",
		"crypto/des",
		"crypto/dsa",
		"crypto/ecdsa",
		"crypto/elliptic",
		"crypto/hmac",
		"crypto/md5",
		"crypto/rand",
		"crypto/rc4",
		"crypto/rsa",
		"crypto/sha1",
		"crypto/sha256",
		"crypto/sha512",
		"crypto/subtle",
		"crypto/tls",
		"crypto/x509",
		"crypto/x509/pkix",
		"database",
		"database/sql",
		"database/sql/driver",
		"debug",
		"debug/dwarf",
		"debug/elf",
		"debug/gosym",
		"debug/macho",
		"debug/pe",
		"encoding",
		"encoding/ascii85",
		"encoding/asn1",
		"encoding/base32",
		"encoding/base64",
		"encoding/binary",
		"encoding/csv",
		"encoding/gob",
		"encoding/hex",
		"encoding/json",
		"encoding/pem",
		"encoding/xml",
		"errors",
		"expvar",
		"flag",
		"fmt",
		"go",
		"go/ast",
		"go/build",
		"go/doc",
		"go/format",
		"go/parser",
		"go/printer",
		"go/scanner",
		"go/token",
		"hash",
		"hash/adler32",
		"hash/crc32",
		"hash/crc64",
		"hash/fnv",
		"html",
		"html/template",
		"image",
		"image/color",
		"image/palette",
		"image/draw",
		"image/gif",
		"image/jpeg",
		"image/png",
		"index",
		"index/suffixarray",
		"io",
		"io/ioutil",
		"log",
		"log/syslog",
		"math",
		"math/big",
		"math/cmplx",
		"math/rand",
		"mime",
		"mime/multipart",
		"net",
		"net/http",
		"net/http/cgi",
		"net/http/cookiejar",
		"net/http/fcgi",
		"net/http/httptest",
		"net/http/httputil",
		"net/http/pprof",
		"net/mail",
		"net/rpc",
		"net/rpc/jsonrpc",
		"net/smtp",
		"net/textproto",
		"net/url",
		"os",
		"os/exec",
		"os/signal",
		"os/user",
		"path",
		"path/filepath",
		"reflect",
		"regexp",
		"regexp/syntax",
		"runtime",
		"runtime/cgo",
		"runtime/debug",
		"runtime/pprof",
		"runtime/race",
		"sort",
		"strconv",
		"strings",
		"sync",
		"sync/atomic",
		"syscall",
		"testing",
		"testing/iotest",
		"testing/quick",
		"text",
		"text/scanner",
		"text/tabwriter",
		"text/template",
		"text/template/parse",
		"time",
		"unicode",
		"unicode/utf16",
		"unicode/utf8",
		"unsafe",
	}
)

func main() {
	allPkgs := flag.Bool("a", false, "Show all packages")
	help := flag.Bool("h", false, "Show this usage")
	flag.Usage = usage
	flag.Parse()

	if *allPkgs {
		showAllPkgs()
		os.Exit(0)
	}

	if *help || len(flag.Args()) != expectedArgsLen {
		flag.Usage()
		os.Exit(0)
	}

	pkg := flag.Args()[0]

	if pkgNotFound(pkg) {
		fmt.Printf("Package %s was not found.\n", pkg)
		showSimilarPkgs(pkg)
		os.Exit(0)
	}

	url := baseUrl + pkg + "/"

	switch runtime.GOOS {
	case "linux":
		exec.Command("xdg-open", url).Start()
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		exec.Command("open", url).Start()
	default:
		fmt.Println("Your PC is not supported.")
	}
}

func pkgNotFound(pkg string) bool {
	for _, p := range pkgs {
		if pkg == p {
			return false
		}
	}
	return true
}

func usage() {
	fmt.Println("Usage: goc [OPTIONS] PACKAGE")
	fmt.Println("Options:")
	flag.PrintDefaults()
}

func showAllPkgs() {
	fmt.Println("Packages:")
	for _, p := range pkgs {
		fmt.Println(p)
	}
}

func showSimilarPkgs(pkg string) {
	sPkgs := make([]string, 0)

	for _, p := range pkgs {
		if strings.Index(p, pkg) > -1 {
			sPkgs = append(sPkgs, p)
		}
	}

	if len(sPkgs) > 0 {
		fmt.Println("Similar packages:")
		for _, p := range sPkgs {
			fmt.Println(p)
		}
	}
}
