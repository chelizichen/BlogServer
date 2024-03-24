<script lang="ts">
export default {
  name: "upload-file",
};
</script>
<template>
  <div>
    <div style="display: flex; align-items: center; justify-content: space-around">
      <h1 style="color: #712424">优品-中原远程包发布与管理中心</h1>
      <input type="file" ref="fileRef" />
    </div>
    <el-card shadow="hover" v-for="(item, key) in uploadList" :key="key">
      <div style="display: flex; align-items: center; justify-content: space-between">
        <div>
          <span style="font-weight: bold"
            >文件名称 : <span style="color: #712424">{{ item.name }}</span></span
          >
          <span> | </span>
          <span style="font-weight: bold">哈希值 : {{ item.hash }}</span>
        </div>
        <div>
          <el-button type="text" @click="getPacakge(item)">下载</el-button>
          <el-button type="text" @click="deletePackage(item)" style="color: red"
            >删除</el-button
          >
        </div>
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import API from "@/api/upload";
import { onMounted, ref } from "vue";
import SparkMD5 from "spark-md5";
const chunkSize = 1024 * 1024 * 20;
const fileRef = ref<HTMLInputElement>();
function initDom() {
  fileRef.value!.onchange = function (e) {
    console.log("this.files", this);

    const file = this.files[0];
    const sliceBuffer = [];
    let sliceSize = file.size;
    while (sliceSize > chunkSize) {
      const blobPart = file.slice(
        sliceBuffer.length * chunkSize,
        (sliceBuffer.length + 1) * chunkSize
      );
      sliceBuffer.push(blobPart);
      sliceSize -= chunkSize;
    }

    if (sliceSize > 0) {
      sliceBuffer.push(file.slice(sliceBuffer.length * chunkSize, file.size));
    }
    let len = 0;
    const fileReader = new FileReader();
    fileReader.onload = function (res) {
      const result = fileReader.result;
      const fileHash = SparkMD5.hashBinary(result);

      API.checkFileChunkState(fileHash)
        .then((res) => {
          let { chunkList, state } = res;
          if (state === 1) {
            alert("已经上传完成");
            return;
          }

          chunkList = chunkList.map((e) => parseInt(e));

          const chunkRequests = [];
          sliceBuffer.forEach((buffer, i) => {
            if (!chunkList.includes(i)) {
              const blob = new File([buffer], `${i}`);
              chunkRequests.push(API.uploadFileChunk(fileHash, blob));
            }
          });
          len = chunkRequests.length;
          return Promise.all(chunkRequests);
        })
        .then((res) => {
          let timer: unknown = setInterval(async () => {
            const resp = await API.megerChunkFile(fileHash, file.name, len);
            if (!resp.code) {
              console.log("resp", resp);
              clearInterval(timer);
              timer = null;
              init();
            }
          }, 2000);
        })
        .then((res) => {
          console.log(res);
        });
    };
    fileReader.onerror = function (err) {
      console.log("报错了", err.target.error);
    };
    fileReader.readAsBinaryString(this.files[0]);
  };
}
const uploadList = ref([]);
function getPacakge(item) {
  const { hash, name } = item;
  const filePath = `/blogserver/uploadPackages/${hash}/${name}`;
  window.open(filePath);
}
async function deletePackage(item) {
  const { hash, name } = item;
  const filePath = `/uploadFile/${hash}`;
  await API.DeletePackage({ dir: filePath });
  await init();
}
async function init() {
  const resp = await API.GetFiles();
  uploadList.value = resp.data;
}

onMounted(() => {
  init();
  initDom();
});
</script>

<style scoped></style>
