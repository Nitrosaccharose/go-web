<script setup>
import request from "/src/request.js"
import {ElMessage} from "element-plus"

let UserList = $ref([]);

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

</script>

<template>
  <div style="height: 100%; display: flex; justify-content: center; align-items: center;">
    <div style="height: 70%; width: 700px; display: flex; flex-direction: column; justify-content: center; align-items: center;">
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
</template>

<style scoped>

</style>