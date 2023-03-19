<script setup>
import request from "/src/request.js"
import {ElMessage} from "element-plus"

let UserList = $ref([]);
let KVTList = $ref([])

const upDataTableData = () => {
  request.get('/user/get/all')
      .then(response => {
        if (response.status === 200) {
          UserList = response.data.data
        } else {
          ElMessage({
            type: 'error',
            message: '获取信息失败！'
          });
        }
      })
      .catch(err => {
        ElMessage({
          type: 'error',
          message: '连接数据库失败！'
        });
      });
}
upDataTableData()
// 批量删除数据的id信息
let multipleSelection = $ref([])

// 激励新增/编辑对话框，为true时显示对话框，反之影藏
let dialogFormVisible = $ref(false)
let drawerVisible = $ref(false)
// 表单对象
let tableForm = $ref({})

//弹窗状态，只有add和edit两种状态，用于表示表单状态
let dialogType = $ref('add')


// 响应新增事件
const handleAdd = () => {
  dialogFormVisible = true
  dialogType = 'add'
  tableForm = {}
}

const handleRandomKVT = () => {
  const generateRandomString = (length) => {
    let result = '';
    const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    for (let i = 0; i < length; i++) {
      result += characters.charAt(Math.floor(Math.random() * characters.length));
    }
    return result;
  };

  const key = generateRandomString(Math.floor(Math.random() * 3) + 4);
  const value = generateRandomString(Math.floor(Math.random() * 3) + 4);
  const time = Math.floor(Math.random() * 20) + 5;

  const datakvt = {
    key,
    value,
    time
  };
  console.log(datakvt)
  request.post('/kvt/post', datakvt).then(response => {

    KVTList.push(datakvt)
    if (response.status === 200) {
      ElMessage({
        type: 'success',
        message: '新增KVT成功！'
      })
    }
  }).catch(err => {
    ElMessage({
      type: 'error',
      message: '新增KVT失败！'
    })
  })

}

// 定义一个函数来进行监控操作
function monitorIKVT() {
  for (let i = 0; i < KVTList.length; i++) {
    if (KVTList[i].time <= 0) {
      KVTList.splice(i, 1);
      i--;
      continue
    }
    request.get('/kvt/get/key' + "?key=" + KVTList[i].key).then(response => {
      if (response.status === 200) {
        KVTList[i].time = response.data.data.Time
      }
    }).catch(err => {
      KVTList.splice(i, 1);
      i--;
    })

  }
}
setInterval(monitorIKVT, 650);

const tableRowClassName = (row) => {
  if (row.row.time <= 2) {
    return 'danger-row'
  }
  if (row.row.time <= 4) {
    return 'warning-row'
  }
  return '';
};


// 提交表单
const dialogConfirm = () => {
  if (dialogType === 'add') {
    console.log(tableForm)
    request.post('/user/post', tableForm).then(response => {
      if (response.status === 200) {
        console.log(response)
        ElMessage({
          type: 'success',
          message: '新增信息成功！'
        })
        dialogFormVisible = false
        upDataTableData()
      }
    }).catch(err => {
      ElMessage({
        type: 'error',
        message: '新增信息失败！'
      })
    })
  } else {
    request.put('/user/put', tableForm).then(response => {
      if (response.status === 200) {
        ElMessage({
          type: 'success',
          message: '修改信息成功！'
        })
        dialogFormVisible = false
        upDataTableData()
      }
    }).catch(err => {
      ElMessage({
        type: 'error',
        message: '修改信息失败！'
      })
    })
  }
}

// 依据行id删除数据的方法
const handleRowDel = (dataobj) => {
  request.delete('/user/delete',  { data: dataobj }).then(response => {
    if (response.status === 200) {
      ElMessage({
        type: 'success',
        message: '删除信息成功！'
      })
      upDataTableData()
    }
  }).catch(err => {
    ElMessage({
      type: 'error',
      message: '删除信息失败！'
    })
  })
}


// 获取勾选数据的对象，并将其id放入multipleSelection
const handleSelectionChange = (val) => {
  multipleSelection = []
  val.forEach(item => {
    multipleSelection.push(item)
  });
}

// 批量删除的方法
const handleDeleteList = () => {
  console.log(multipleSelection)
  multipleSelection.forEach(dataobj => {
    handleRowDel(dataobj)
  })
  multipleSelection = []
  upDataTableData()
}

// 编辑已有数据的方法
const handleEdit = (row) => {
  dialogFormVisible = true
  dialogType = 'edit'
  tableForm = {...row}
}

let technologyList = [
  {
    "technology": "Vue",
    "introduce": "Vue是一套用于构建用户界面的渐进式框架，具有轻量、高效、易用等特点。Vue的核心库只关注视图层，非常容易上手，可以与其它第三方库或既有项目整合。"
  },
  {
    "technology": "Vite",
    "introduce": "Vite是一款基于浏览器原生 ES 模块构建的前端构建工具，开发模式下利用浏览器原生 ES 模块直接导入文件的方式来提高构建速度和开发效率。"
  },
  {
    "technology": "Vue-router",
    "introduce": "Vue-router是Vue.js官方的路由管理器，用于创建单页面应用。它通过URL映射到组件，实现了前端路由和组件的耦合。"
  },
  {
    "technology": "ElementPlus",
    "introduce": "ElementPlus是一款基于Vue 3.0的UI组件库，包含丰富的组件和强大的功能。它具有简单易用、高效稳定等特点，可以快速搭建优秀的用户界面。"
  },
  {
    "technology": "Axios",
    "introduce": "Axios是一个基于Promise的HTTP客户端，可以用于浏览器和Node.js环境中，支持请求拦截、响应拦截、请求取消等特性，提供了一种优雅、简洁的方式来处理HTTP请求。"
  },
  {
    "technology": "Go",
    "introduce": "Go是一种高效、可靠的编程语言，具有良好的并发编程和内存管理机制，被广泛应用于网络编程、系统编程、云计算、人工智能等领域。"
  },
  {
    "technology": "gin",
    "introduce": "Gin是一个使用Go语言编写的高性能Web框架，具有轻量级、高性能、易用等特点，被广泛应用于RESTful API的开发。"
  },
  {
    "technology": "gorm",
    "introduce": "GORM是一个使用Go语言编写的ORM框架，支持多种数据库，包括MySQL、PostgreSQL、SQLite等，具有简单易用、功能丰富等特点。"
  },
  {
    "technology": "go-redis",
    "introduce": "go-redis是Go语言的Redis客户端库，提供了完整的Redis命令封装和一些额外的功能，具有高性能、易用等特点。"
  },
  {
    "technology": "Docker",
    "introduce": "Docker是一款开源的容器化平台，可以将应用程序及其依赖项打包到一个可移植的容器中，从而实现应用程序的快速部署、跨平台移植等优势。Docker还支持镜像管理、容器网络、存储卷等功能，使得应用程序在不同环境中运行更加方便。"
  },
  {
    "technology": "Redis",
    "introduce": "Redis是一个高性能的键值存储系统，支持多种数据结构，包括字符串、哈希、列表、集合、有序集合等。Redis被广泛应用于缓存、队列、计数器、分布式锁等场景。"
  },
  {
    "technology": "MySQL",
    "introduce": "MySQL是一款开源的关系型数据库管理系统，具有成本低廉、易于使用、稳定性高等特点，被广泛应用于Web应用程序、数据仓库、企业应用等领域。"
  }
]
const handleOpenDrawer = () => {
  drawerVisible = true
}
</script>

<template>
  <div style="height: 100%; display: flex; flex-direction: column; justify-content: center; align-items: center;">
    <div style="margin-top: 2%; margin-bottom: 2%;">
      <div style="font-size: 28pt;">基于Vue + Gin + Redis + MySQL简单GoWeb应用示例</div>
    </div>
    <div style="height: 500px; display: flex; justify-content: center; align-items: center;">
     <div style="height: 100%; width: 700px; display: flex; flex-direction: column; justify-content: flex-start; align-items: center;">
       <div style="width:100%; display: flex; justify-content: space-between;">
         <el-button type="primary" @click="handleAdd">新增</el-button>
         <el-button type="danger" @click="handleDeleteList">删除</el-button>
       </div>
       <el-table :data="UserList" style="width: 100%" @selection-change="handleSelectionChange">
         <el-table-column type="selection" width="40"/>
         <el-table-column prop="user_account" label="账号" width="180"/>
         <el-table-column prop="user_name" label="昵称" width="180"/>
         <el-table-column prop="user_password" label="密码"/>
         <el-table-column label="操作">
           <template #default="scope">
             <div style="display: flex; flex-wrap: nowrap;">
               <el-button plain type="primary" size="small" @click="handleEdit(scope.row)">编辑</el-button>
               <el-button plain type="danger" size="small" @click="handleRowDel(scope.row)">删除</el-button>
             </div>
           </template>

         </el-table-column>
       </el-table>
     </div>
     <div style="height: 100%; width: 500px; display: flex; flex-direction: column; justify-content: flex-start; align-items: center; margin-left: 50px;">
       <div style="width:100%; display: flex; justify-content: center;">
         <el-button color="#626aef" @click="handleRandomKVT">随机添加 (Key,Value,Time)</el-button>
       </div>
       <el-table :data="KVTList" style="width: 100%" :row-class-name="tableRowClassName">
         <el-table-column prop="key" label="Key" width="180"/>
         <el-table-column prop="value" label="Value" width="180"/>
         <el-table-column prop="time" label="Time"/>
       </el-table>
     </div>
   </div>
    <div style="margin-top: 2%; margin-bottom: 2%;">
      <el-button color="#DCDCDC" size="large" @click="handleOpenDrawer">技术选型</el-button>
    </div>
  </div>

  <el-dialog v-model="dialogFormVisible" :title="dialogType === 'add' ? '新增' : '编辑'" width="40%">
    <el-form :model="tableForm" label-width="auto">
      <el-form-item label="账号" prop="user_account">
        <el-input v-model="tableForm.user_account" autocomplete="off"/>
      </el-form-item>
      <el-form-item label="昵称" prop="user_name">
        <el-input v-model="tableForm.user_name" autocomplete="off"/>
      </el-form-item>
      <el-form-item label="密码" prop="user_password">
        <el-input v-model="tableForm.user_password" autocomplete="off"/>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button type="primary" @click="dialogConfirm">确认</el-button>
      </span>
    </template>
  </el-dialog>

  <el-drawer
      v-model="drawerVisible"
      title="技术选型"
      :direction="'ltr'"
      size="50%"

  >
    <div style="display: flex; justify-content: center; align-items: center;">
      <el-table :data="technologyList" style="width: 100%">
        <el-table-column prop="technology" label="语言框架"/>
        <el-table-column prop="introduce" label="简介"/>
      </el-table>
    </div>

  </el-drawer>
</template>

<style>
.el-table .warning-row {
  --el-table-tr-bg-color: var(--el-color-warning-light-9);
}
.el-table .danger-row {
  --el-table-tr-bg-color: var(--el-color-danger-light-9);
}
</style>