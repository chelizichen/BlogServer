package service

import (
	"Simp/servers/BlogServer/utils"
	handlers "Simp/src/http"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Cors(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}

func UploadService(ctx *handlers.SimpHttpServerCtx, pre string) {
	E := ctx.Engine
	G := E.Group(pre)
	G.Use(Cors)
	G.GET("/checkChunk", func(c *gin.Context) {
		hash := c.Query("hash")
		hashPath := fmt.Sprintf("./uploadFile/%s", hash)
		hashPath = utils.JoinPathUtil(hashPath)
		chunkList := []string{}
		isExistPath, err := PathExists(hashPath)
		if err != nil {
			fmt.Println("获取hash路径错误", err)
		}

		if isExistPath {
			files, err := os.ReadDir(hashPath)
			state := 0
			if err != nil {
				fmt.Println("文件读取错误", err)
			}
			for _, f := range files {
				fileName := f.Name()
				chunkList = append(chunkList, fileName)
				fileBaseName := strings.Split(fileName, ".")[0]
				if fileBaseName == hash {
					state = 1
				}
			}

			c.JSON(200, gin.H{
				"state":     state,
				"chunkList": chunkList,
			})
		} else {
			c.JSON(200, gin.H{
				"state":     0,
				"chunkList": chunkList,
			})
		}
	})

	G.POST("/uploadChunk", func(c *gin.Context) {
		fileHash := c.PostForm("hash")
		file, err := c.FormFile("file")
		hashPath := fmt.Sprintf("./uploadFile/%s", fileHash)
		hashPath = utils.JoinPathUtil(hashPath)
		if err != nil {
			fmt.Println("获取上传文件失败", err)
		}

		isExistPath, err := PathExists(hashPath)
		if err != nil {
			fmt.Println("获取hash路径错误", err)
		}

		if !isExistPath {
			os.Mkdir(hashPath, os.ModePerm)
		}
		s := fmt.Sprintf("./uploadFile/%s/%s", fileHash, file.Filename)
		fmt.Println("SaveUploadedFile before", s)
		s = utils.JoinPathUtil(s)
		fmt.Println("SaveUploadedFile after", s)
		// go func() {
		go func() {
			go c.SaveUploadedFile(file, s)
		}()
		if err != nil {
			c.String(400, "0")
			fmt.Println(err)
		} else {
			chunkList := []string{}
			files, err := os.ReadDir(hashPath)
			if err != nil {
				fmt.Println("文件读取错误", err)
			}
			for _, f := range files {
				fileName := f.Name()

				if f.Name() == ".DS_Store" {
					continue
				}
				chunkList = append(chunkList, fileName)
			}

			c.JSON(200, gin.H{
				"chunkList": chunkList,
			})
		}
		// }()
	})

	G.GET("/megerChunk", func(c *gin.Context) {
		hash := c.Query("hash")
		fileName := c.Query("fileName")
		size := c.Query("size") // 传过来的需要和文件中的一样
		Size, err := strconv.Atoi(size)
		hashPath := fmt.Sprintf("./uploadFile/%s", hash)
		hashPath = utils.JoinPathUtil(hashPath)
		isExistPath, err := PathExists(hashPath)
		if err != nil {
			fmt.Println("获取hash路径错误", err)
			c.AbortWithStatusJSON(200, err.Error())
			return
		}

		if !isExistPath {
			c.AbortWithStatusJSON(400, gin.H{
				"message": "文件夹不存在",
				"code":    "-1",
			})
			return
		}
		isExistFile, err := PathExists(hashPath + "/" + fileName)
		if err != nil {
			fmt.Println("获取hash路径文件错误", err)
			c.AbortWithStatusJSON(400, gin.H{
				"message": "获取hash路径文件错误" + err.Error(),
				"code":    "-1",
			})
			return
		}
		fmt.Println("文件是否存在", isExistFile)
		if isExistFile {
			c.JSON(200, gin.H{
				"fileUrl": fmt.Sprintf("/uploadFile/%s/%s", hash, fileName),
				"code":    "0",
				"data":    "文件已存在",
			})
			return
		}

		files, err := os.ReadDir(hashPath)
		currentSize := len(files)
		if Size != currentSize {
			c.JSON(500, gin.H{
				"message": "还没打包好",
				"code":    -1,
			})
		}
		if err != nil {
			fmt.Println("合并文件读取失败", err)
		}
		complateFile, err := os.Create(hashPath + "/" + fileName)
		if err != nil {
			fmt.Println("Create", err)
			c.AbortWithStatusJSON(400, gin.H{
				"message": "CreateError" + err.Error(),
				"code":    "-1",
			})
			return
		}
		for _, f := range files {
			//.DS_Store
			// file, err := os.Open(hashPath + "/" + f.Name())
			// if err != nil {
			// 	fmt.Println("文件打开错误", err)
			// }

			// if f.Name() == ".DS_Store" {
			// continue
			// }

			fileBuffer, err := os.ReadFile(hashPath + "/" + f.Name())
			if err != nil {
				fmt.Println("文件打开错误", err)
			}
			complateFile.Write(fileBuffer)
		}
		c.JSON(200, gin.H{
			"fileUrl": fmt.Sprintf("/uploadFile/%s/%s", hash, fileName),
			"code":    0,
		})
		defer complateFile.Close()

	})

	G.GET("/getFiles", func(ctx *gin.Context) {
		var uploadPath string = "uploadFile"
		var staticPath string = utils.JoinPathUtil(uploadPath)
		fmt.Println("staticPath", staticPath)
		de, err := os.ReadDir(staticPath)
		dirString := make([]map[string]interface{}, 0)
		if err != nil {
			fmt.Println("Read Dir Error", err.Error())
			ctx.AbortWithStatusJSON(200, gin.H{
				"data": dirString,
			})
			return
		}
		for _, v := range de {
			m := make(map[string]interface{})
			if v.IsDir() {
				m["hash"] = v.Name()
				s := filepath.Join("uploadFile", v.Name())
				s = utils.JoinPathUtil(s)
				fmt.Println("dir || ", s)
				tgz, err := os.ReadDir(s)
				if err != nil {
					fmt.Println("error to getServer Dir", err.Error())
				}
				for _, v := range tgz {
					fmt.Println("tgz", v.Name())
					// if strings.HasSuffix(v.Name(), ".tgz") || strings.HasSuffix(v.Name(), ".gz") {
					m["name"] = v.Name()
					// }
				}
				dirString = append(dirString, m)
			}
		}
		ctx.JSON(200, gin.H{
			"data": dirString,
		})
	})

	G.GET("/deletePackage", func(ctx *gin.Context) {
		var dir = ctx.Query("dir")
		staticPath := utils.JoinPathUtil(dir)
		fmt.Println("delete staticPath", staticPath)
		err := DeleteDirectory(staticPath)
		if err != nil {
			fmt.Println("Error to remove All", err.Error())
		}
		ctx.JSON(http.StatusOK, nil)
	})
}

func CreateUploadPath() {
	s := utils.JoinPathUtil("uploadFile")
	b, err := PathExists(s)
	if err != nil {
		fmt.Println("CreateUploadPath | Error", err.Error())
		return
	}
	if !b {
		os.Mkdir(s, os.ModePerm)
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func DeleteDirectory(path string) error {
	// 获取目录中的文件和子目录
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		var currentPath string
		if file.IsDir() {
			// 如果是子目录，递归调用 DeleteDirectory 函数
			currentPath = fmt.Sprintf("%s/%s", path, file.Name())
			if err := DeleteDirectory(currentPath); err != nil {
				fmt.Printf("Error deleting directory: %v\n", err)
			}
		} else {
			// 如果是文件，直接删除文件
			currentPath = fmt.Sprintf("%s/%s", path, file.Name())
			err := os.Remove(currentPath)
			if err != nil {
				fmt.Printf("Error deleting file: %v\n", err)
			}
		}
	}

	// 删除目录本身
	err = os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}
