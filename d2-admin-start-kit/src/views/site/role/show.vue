<template>
  <d2-container>
    <template slot="header">角色编辑</template>
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
        <el-form-item label="角色" prop="name">
          <el-input
            v-model="ruleForm.name"
            placeholder="请输入路由名称"
          ></el-input>
        </el-form-item>

        <el-form-item label="路由简介" prop="desc">
          <el-input
            v-model="ruleForm.desc"
            placeholder="请输入路由简介"
          ></el-input>
        </el-form-item>

        <el-form-item label="排序" prop="order">
          <el-input-number
            v-model="ruleForm.order"
            :min="0"
            :max="10000"
            label="请选择排序  "
          ></el-input-number>
        </el-form-item>

        <el-form-item label="角色权限" prop="permissions">
          <el-cascader
            placeholder="试试搜索：指南"
             v-model="ruleForm.permissions"
            :options="permissionTree"
            :props="{ multiple: true, checkStrictly: true }"
            filterable
          ></el-cascader>
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

<script>
import api from "@/api";
import { mapState, mapActions } from "vuex";
export default {
  name: "role.show",
  data() {
    return {
      id: 0,
      edit: false,
      permissionTree: [],
      ruleForm: {
        name: "",
        desc: "",
        order: 0,
        permissions: [],
      },
      rules: {
        name: [
          { required: true, message: "请输入角色名称", trigger: "blur" },
          {
            min: 2,
            max: 10,
            message: "长度在 2 到 10 个字符",
            trigger: "blur",
          },
        ],
        desc: [
          { required: true, message: "请输入角色简介", trigger: "blur" },
          {
            min: 4,
            max: 20,
            message: "长度在 4 到 20 个字符",
            trigger: "blur",
          },
        ],
      },
    };
  },
  computed: {
    ...mapState("d2admin/page", [
      "opened",
      "current", //用户获取当前页面的地址，用于关闭
    ]),
  },
  mounted() {
    this.id = this.$route.params.id;
    this.getPermissionTree();
    if (this.id > 0) {
      this.edit = true;
      this.show();
    }
  },
  methods: {
    ...mapActions("d2admin/page", ["close"]),
    submitForm(formName) {
      if (this.edit) {
        this.rules.Password = [];
      }
      this.$refs[formName].validate((valid) => {
        if (valid) {
          if (this.edit) {
            this.update();
          } else {
            this.store();
          }
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
    async getPermissionTree() {
      const res = await api.SYS_PERMISSION_TREE();
      res.shift();
      this.permissionTree = res;
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    async show() {
      const res = await api.SYS_ROLE_INFO(this.id);
      this.ruleForm = res;
    },
    async update() {
      const res = await api.SYS_ROLE_UPDATE(this.id, this.ruleForm);
      this.ruleForm = res;
      this.$notify({
        title: "成功",
        message: "更新成功",
        type: "success",
        duration: 2000,
      });
    },
    async store() {
      await api.SYS_ROLE_STORE(this.ruleForm);
      this.$notify({
        title: "成功",
        message: "新增成功",
        type: "success",
        duration: 2000,
      });
      setTimeout(() => {
        let tagName = this.current;
        this.close({ tagName });
      }, 2000);
    },
  },
};
</script>
