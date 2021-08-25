<template>
  <d2-container>
    <template slot="header">用户编辑</template>
    <el-card class="box-card">
      <el-form
        :model="ruleForm"
        :rules="rules"
        ref="ruleForm"
        label-width="100px"
        class="demo-ruleForm"
        size="small"
        style="width: 50%; margin: auto"
      >
        <el-form-item label="头像" prop="Avatar" style="text-align: center">
          <el-upload
            class="avatar-uploader"
            action="https://jsonplaceholder.typicode.com/posts/"
            :show-file-list="false"
          >
            <img v-if="ruleForm.Avatar" :src="ruleForm.Avatar" class="avatar" />
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>
        </el-form-item>

        <el-form-item label="用户名称" prop="Name">
          <el-input
            v-model="ruleForm.Name"
            placeholder="请输入用户名称"
          ></el-input>
        </el-form-item>

        <el-form-item label="邮箱" prop="Email">
          <el-input
            v-model="ruleForm.Email"
            placeholder="请输入用户邮箱"
          ></el-input>
        </el-form-item>

        <el-form-item label="用户密码" prop="Password" v-show="!edit">
          <el-input
            v-model="ruleForm.Password"
            placeholder="请输入密码"
            show-password
          ></el-input>
        </el-form-item>

        <el-form-item label="用户状态" prop="Status">
          <el-radio v-model="ruleForm.Status" :label="0">正常</el-radio>
          <el-radio v-model="ruleForm.Status" :label="1">禁用</el-radio>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="submitForm('ruleForm')"
            >保存</el-button
          >
          <el-button @click="resetForm('ruleForm')">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </d2-container>
</template>
<style>
.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-uploader .el-upload:hover {
  border-color: #409eff;
}
.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}
.avatar {
  width: 178px;
  height: 178px;
  display: block;
}
</style>

<script>
import api from "@/api";
export default {
  name: "user.show",
  data() {
    // 邮箱
    var checkEmail = (rule, value, callback) => {
      const mailReg = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+/;
      if (!value) {
        return callback(new Error("邮箱不能为空"));
      }
      setTimeout(() => {
        if (mailReg.test(value)) {
          callback();
        } else {
          callback(new Error("请输入正确的邮箱格式"));
        }
      }, 100);
    };
    return {
      id: 0,
      edit: false,
      ruleForm: {
        Name: "",
        Avatar: "",
        Password: "",
        Status: "",
        Email: "",
        CreatedAt: "",
        updated_at: "",
      },
      rules: {
        Name: [
          { required: true, message: "请输入用户名称", trigger: "blur" },
          { min: 2, max: 10, message: "长度在 2 到 10 个字符", trigger: "blur" },
        ],
        Avatar: [
          { required: true, message: "请上传用户头像", trigger: "blur" },
          { min: 2, max: 1000, message: "长度在 2 到 1000 个字符", trigger: "blur" },
        ],
        Password: [
          { required: true, message: "请输入用户密码", trigger: "blur" },
          {
            min: 8,
            max: 20,
            message: "长度在 8 到 20 个字符",
            trigger: "blur",
          },
        ],
        Status: [
          { required: true, message: "请选择用户状态", trigger: "blur" },
        ],
        Email: [{ validator: checkEmail, trigger: "blur" }],
      },
    };
  },
  mounted() {
    this.id = this.$route.params.id;
    if (this.id > 0) {
      this.edit = true;
      this.show();
    }
  },
  methods: {
    submitForm(formName) {
      if(this.edit){
        this.rules.Password=[]
      }
      this.$refs[formName].validate((valid) => {
        if (valid) {
          if (this.edit){
            this.update()
          }
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    async show() {
      const res = await api.SYS_USER_INFO(this.id);
      console.log(res);
      this.ruleForm = res;
    },
    async update() {
      const res = await api.SYS_USER_UPDATE(this.id,this.ruleForm);
      console.log(res);
      this.ruleForm = res;
    }, 
    async store() {
      const res = await api.SYS_USER_INFO(this.id);
      console.log(res);
      this.ruleForm = res;
    },
  },
};
</script>
