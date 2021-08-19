<template>
  <d2-container>
    <template slot="header">用户管理</template>
    <template>
      <el-table :data="tableData" border style="width: 100%">
        <el-table-column label="日期">
          <template slot-scope="scope">
            <i class="el-icon-time"></i>
            <span style="margin-left: 10px">{{ scope.row.CreatedAt }}</span>
          </template>
        </el-table-column>
        <el-table-column label="姓名">
          <template slot-scope="scope">
            <el-popover trigger="hover" placement="top">
              <p>姓名: {{ scope.row.Name }}</p>
              <p>邮箱: {{ scope.row.Email }}</p>
              <div slot="reference" class="name-wrapper">
                <el-tag size="medium">{{ scope.row.Name }}</el-tag>
              </div>
            </el-popover>
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
    </template></d2-container
  >
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
    };
  },
  mounted() {
    this.getList(1)
  },
  methods: {
    async getList(page) {
      const res = await api.SYS_USER_LIST(1);
      console.log(res);
      this.tableData = res.users;
    },
  },
};
</script>
