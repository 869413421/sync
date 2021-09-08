<template>
  <d2-container>
    <template slot="header">权限编辑</template>
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
        <el-form-item label="路由" prop="v_1">
          <el-input v-model="ruleForm.v_1" placeholder="请输入路由"></el-input>
        </el-form-item>

        <el-form-item label="请求方式" prop="v_2">
          <el-select v-model="ruleForm.v_2" placeholder="请选择">
            <el-option
              v-for="item in requestMethods"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            >
            </el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="路由名称" prop="name">
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
  name: "permssion.show",
  data() {
    return {
      id: 0,
      edit: false,
      permssionTree: [],
      requestMethods: [
        {
          value: "GET",
          label: "GET",
        },
        {
          value: "POST",
          label: "POST",
        },
        {
          value: "PUT",
          label: "PUT",
        },
        {
          value: "DELETE",
          label: "DELETE",
        },
        {
          value: "HEAD",
          label: "HEAD",
        },
        {
          value: "OPTIONS",
          label: "OPTIONS",
        },
      ],
      ruleForm: {
        v_1: "",
        v_2: "",
        name: "",
        desc: "",
        ptype: "p",
      },
      rules: {
        name: [
          { required: true, message: "请输入路由名称", trigger: "blur" },
          {
            min: 2,
            max: 10,
            message: "长度在 2 到 10 个字符",
            trigger: "blur",
          },
        ],
        desc: [
          { required: true, message: "请输入路由简介", trigger: "blur" },
          {
            min: 4,
            max: 20,
            message: "长度在 4 到 20 个字符",
            trigger: "blur",
          },
        ],
        v_2: [{ required: true, message: "请选择请求尝试", trigger: "change" }],
        v_1: [
          { required: true, message: "请输入路由", trigger: "blur" },
          {
            min: 2,
            max: 50,
            message: "长度在 2 到 10 个字符",
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
    this.getPermssionTree();
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
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    async show() {
      const res = await api.SYS_CASBIN_INFO(this.id);
      this.ruleForm = res;
    },
    async getPermssionTree() {
      const res = await api.SYS_CASBIN_TREE();
      this.permssionTree = res;
    },
    async update() {
      const res = await api.SYS_CASBIN_UPDATE(this.id, this.ruleForm);
      this.ruleForm = res;
      this.$notify({
        title: "成功",
        message: "更新成功",
        type: "success",
        duration: 2000,
      });
    },
    async store() {
      const res = await api.SYS_CASBIN_STORE(this.ruleForm);
      this.$notify({
        title: "成功",
        message: "新增成功",
        type: "success",
        duration: 2000,
      });
      setTimeout(() => {
        let tagName = this.current;
        console.log(tagName);
        this.close({ tagName });
      }, 2000);
    },
  },
};
</script>
