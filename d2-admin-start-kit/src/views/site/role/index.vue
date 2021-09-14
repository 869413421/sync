<template>
  <d2-container>
    <template slot="header"
      >角色管理
      <el-button type="primary" style="float: right" @click="showEdit"
        >新增角色</el-button
      >
    </template>
    <template>
      <el-table :data="tableData" border style="width: 100%">
        <el-table-column label="角色">
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{ scope.row.name }}</span>
          </template>
        </el-table-column>
        <el-table-column label="角色描述">
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{ scope.row.desc }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button size="mini" @click="handleEdit(scope.$index, scope.row)"
              >编辑</el-button
            >
            <el-button
              size="mini"
              type="danger"
              @click="handleDelete(scope.$index, scope.row)"
              >删除</el-button
            >
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        style="margin-top: 1%"
        background
        layout="prev, pager, next"
        :page-size="10"
        :current-page="currentPage"
        :total="total"
        @current-change="getList"
      >
      </el-pagination>
    </template>
  </d2-container>
</template>

<script>
import api from "@/api";
export default {
  name: "role",
  data() {
    return {
      tableData: [],
      currentPage: 1,
      total: 0,
    };
  },
  mounted() {
    this.getList(1);
  },
  methods: {
    async getList(page) {
      const res = await api.SYS_ROLE_LIST(page);
      this.tableData = res.roles;
      this.currentPage = res.PagerData.Current.Number;
      this.total = res.PagerData.TotalCount;
    },
    handleEdit(index, row) {
      this.$router.push({ name: "role.show", params: { id: row.id } });
    },
    showEdit() {
      this.$router.push({ name: "role.show", params: { id: 0 } });
    },
    async handleDelete(index, row) {
      await api.SYS_ROLE_DELETE(row.id);
      this.$notify({
        title: "成功",
        message: "删除成功",
        type: "success",
        duration: 2000,
      });
      setTimeout(() => {
        this.$router.go(0);
      }, 2000);
    },
  },
};
</script>
