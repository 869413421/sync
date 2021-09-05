<template>
  <d2-container>
    <template slot="header"
      >用户管理
      <el-button type="primary" style="float: right" @click="showEdit">新增用户</el-button>
    </template>
    <template>
      <el-table :data="tableData" border style="width: 100%">
        <el-table-column label="头像">
          <template slot-scope="scope">
            <el-avatar size="small" :src="scope.row.avatar"></el-avatar>
          </template>
        </el-table-column>
        <el-table-column label="姓名">
          <template slot-scope="scope">
            <el-popover trigger="hover" placement="top">
              <p>姓名: {{ scope.row.name }}</p>
              <p>邮箱: {{ scope.row.email }}</p>
              <div slot="reference" class="name-wrapper">
                <el-tag size="medium">{{ scope.row.name }}</el-tag>
              </div>
            </el-popover>
          </template>
        </el-table-column>
        <el-table-column label="创建时间">
          <template slot-scope="scope">
            <i class="el-icon-time"></i>
            <span style="margin-left: 10px">{{ scope.row.created_at }}</span>
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
  </d2-container>
</template>

<script>
import api from "@/api";
export default {
  name: "user",
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
      const res = await api.SYS_USER_LIST(page);
      console.log(res);
      this.tableData = res.users;
      this.currentPage = res.PagerData.Current.Number;
      this.total = res.PagerData.TotalCount;
    },
    handleEdit(index, row) {
       this.$router.push({ name: "user.show", params: { id: row.id } });
    },
    showEdit() {
      this.$router.push({ name: "user.show", params: { id: 0 } });
    },
  },
};
</script>
