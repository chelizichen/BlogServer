export const toolbarConfig = {};
export const editorConfig = {
  placeholder: "请输入内容...",
  MENU_CONF: {
    uploadImage: {
      server: "/blogserver/blogImg",
      // file 文件的字段名
      fieldName: "file",
      maxFileSize: 2 * 1024 * 1024,
    },
  },
};