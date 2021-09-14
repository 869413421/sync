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
        <el-form-item label="路由" prop="url">
          <el-input v-model="ruleForm.url" placeholder="请输入路由"></el-input>
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

        <el-form-item label="请求方式" prop="method">
          <el-select v-model="ruleForm.method" placeholder="请选择">
            <el-option
              v-for="item in requestMethods"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            >
            </el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="父级路由" prop="parent_id">
          <el-cascader
            v-model="ruleForm.parent_id"
            :options="permissionTree"
            :show-all-levels="true"
            :props="{ checkStrictly: true }"
            clearable
          ></el-cascader>
        </el-form-item>

        <el-form-item label="排序" prop="order">
          <el-input-number
            v-model="ruleForm.order"
            :min="0"
            :max="10000"
            label="请选择排序  "
          ></el-input-number>
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
  name: "permission.show",
  data() {
    // 父级路由
    var checkParentId = (rule, value, callback) => {
      setTimeout(() => {
        console.log(value);
        let parent_id = value[value.length - 1];
        if (this.ruleForm.id > 0 && this.ruleForm.id == parent_id) {
          callback(new Error("不允许选中自己为父级菜单"));
        } else {
          this.ruleForm.parent_id = parent_id;
          callback();
        }
      }, 100);
    };
    return {
      id: 0,
      edit: false,
      permissionTree: [],
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
        url: "",
        name: "",
        method: "",
        desc: "",
        order: 0,
        parent_id: [0],
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
        method: [
          { required: true, message: "请选择请求方式", trigger: "change" },
        ],
        url: [
          { required: true, message: "请输入路由", trigger: "blur" },
          {
            min: 2,
            max: 50,
            message: "长度在 2 到 10 个字符",
            trigger: "blur",
          },
        ],
        parent_id: [{ validator: checkParentId, trigger: "blur" }],
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
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    async show() {
      const res = await api.SYS_PERMISSION_INFO(this.id);
      this.ruleForm = res;
      if (res.parent_id == "") {
        this.ruleForm.parent_id = [res.parent_id];
      } else {
        this.ruleForm.parent_id = res.parent_ids
          .split(",")
          .map(function (data) {
            return +data;
          });

        console.log(res.parent_ids.split(","));
      }
    },
    async getPermissionTree() {
      const res = await api.SYS_PERMISSION_TREE();
      this.permissionTree = res;
    },
    async update() {
      const res = await api.SYS_PERMISSION_UPDATE(this.id, this.ruleForm);
      this.ruleForm = res;
      this.ruleForm.parent_id = [this.ruleForm.parent_id];
      this.$notify({
        title: "成功",
        message: "更新成功",
        type: "success",
        duration: 2000,
      });
    },
    async store() {
      const res = await api.SYS_PERMISSION_STORE(this.ruleForm);
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
