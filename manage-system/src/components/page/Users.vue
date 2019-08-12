<template>
    <div class="table">
        <div class="crumbs">
            <el-breadcrumb separator="/">
                <el-breadcrumb-item><i class="el-icon-lx-cascades"></i> 用户列表</el-breadcrumb-item>
            </el-breadcrumb>
        </div>
        <div class="container">
            <div class="handle-box">
                <el-button type="primary" icon="el-icon-delete" class="handle-del mr10" @click="delAll">批量删除</el-button>
                <!--                <el-select v-model="select_cate" placeholder="筛选省份" class="handle-select mr10">-->
                <!--                    <el-option key="1" label="广东省" value="广东省"></el-option>-->
                <!--                    <el-option key="2" label="湖南省" value="湖南省"></el-option>-->
                <!--                </el-select>-->
                <el-input v-model="select_word" placeholder="用户名" class="handle-input mr10"></el-input>
                <el-button type="primary" icon="el-icon-search" @click="search">搜索</el-button>
                <el-button type="primary" icon="el-icon-news" @click="addUser">新增</el-button>
            </div>
            <el-table :data="data" border class="table" ref="userTable" @selection-change="handleSelectionChange">
                <el-table-column type="selection" width="55" align="center"></el-table-column>
                <el-table-column prop="Id" label="id" sortable width="150"></el-table-column>
                <el-table-column prop="Username" label="用户名" width="150"></el-table-column>
                <el-table-column prop="Password" label="密码" width="150"></el-table-column>
                <el-table-column prop="Name" label="名字" :formatter="formatter"></el-table-column>
                <el-table-column label="操作" width="180" align="center">
                    <template slot-scope="scope">
                        <el-button type="text" v-if="scope.row.Id!=1" icon="el-icon-edit"
                                   @click="handleEdit(scope.$index, scope.row)">编辑
                        </el-button>
                        <el-button type="text" icon="el-icon-delete" class="red"
                                   @click="handleDelete(scope.$index, scope.row)">删除
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="pagination">
                <el-pagination background @current-change="handleCurrentChange" layout="prev, pager, next"
                               :total="total">
                </el-pagination>
            </div>
        </div>

        <!-- 编辑弹出框 -->
        <el-dialog title="编辑" :visible.sync="editVisible" width="30%">
            <el-form ref="form" :model="form" label-width="80px">
                <el-form-item label="用户名">
                    <el-input v-model="form.Username"></el-input>
                </el-form-item>
                <el-form-item label="密码">
                    <el-input v-model="form.Password"></el-input>
                </el-form-item>
                <el-form-item label="姓名">
                    <el-input v-model="form.Name"></el-input>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="editVisible = false">取 消</el-button>
                <el-button type="primary" @click="saveEdit">确 定</el-button>
            </span>
        </el-dialog>

        <!-- 删除提示框 -->
        <el-dialog title="提示" :visible.sync="delVisible" width="300px" center>
            <div class="del-dialog-cnt">删除不可恢复，是否确定删除？</div>
            <span slot="footer" class="dialog-footer">
                <el-button @click="delVisible = false">取 消</el-button>
                <el-button type="primary" @click="deleteRow">确 定</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>
    import {addUser, deleteUser, updateUser, userData} from '../../api/index';

    export default {
        name: 'userTable',
        data() {
            return {
                tableData: [],
                cur_page: 1,
                type: 0,
                total: 100,
                limit: 10,
                multipleSelection: [],
                select_cate: '',
                select_word: '',
                del_list: [],
                is_search: false,
                editVisible: false,
                delVisible: false,
                form: {
                    name: '',
                    date: '',
                    address: ''
                },
                idx: -1,
                id: -1
            }
        },
        created() {
            this.getData();
        },
        computed: {
            data() {
                if (this.tableData) {
                    return this.tableData.filter((d) => {

                        let is_del = false;
                        // for (let i = 0; i < this.del_list.length; i++) {
                        //     if (d.name === this.del_list[i].name) {
                        //         is_del = true;
                        //         break;
                        //     }
                        // }
                        if (!is_del) {
                            // if (d.address.indexOf(this.select_cate) > -1 &&
                            //     (d.name.indexOf(this.select_word) > -1 ||
                            //         d.address.indexOf(this.select_word) > -1)
                            // ) {
                            //     return d;
                            // }
                        }
                        return d;
                    })
                }
            }
        },
        methods: {
            // 分页导航
            handleCurrentChange(val) {
                this.cur_page = val;
                this.getData();
            },
            // 获取 easy-mock 的模拟数据
            getData() {
                userData({
                    page: this.cur_page,
                    limit: 10,
                    select_word: this.select_word
                }).then((res) => {
                    this.tableData = res.data;
                    this.total = res.total;
                })
            },
            search() {
                this.getData();
            },
            addUser() {
                this.form = {
                    Id: 0,
                    Name: '',
                    Password: '',
                    Username: ''
                };
                this.editVisible = true;
                this.type = 1;
            },
            formatter(row, column) {
                return row.Name;
            },
            filterTag(value, row) {
                return row.tag === value;
            },
            handleEdit(index, row) {
                this.type = 0;
                this.idx = index;
                this.Id = row.Id;
                this.form = {
                    Id: row.Id,
                    Name: row.Name,
                    Password: row.Password,
                    Username: row.Username
                };
                this.editVisible = true;
            },
            handleDelete(index, row) {
                this.idx = index;
                this.id = row.Id;
                this.delVisible = true;
            },
            delAll() {
                const length = this.multipleSelection.length;
                let str = '';
                let ids = [];
                this.del_list = this.del_list.concat(this.multipleSelection);
                for (let i = 0; i < length; i++) {
                    str += this.multipleSelection[i].Name + ' ';
                    ids.push(this.multipleSelection[i].Id)
                }
                deleteUser({
                    ids: ids,
                }).then((res) => {
                    if (res.code == 0) {
                        this.$message.success('删除了' + str);
                        this.multipleSelection = [];
                        this.getData();
                    } else {
                        this.$message.error(res.message);
                    }
                });
            },
            handleSelectionChange(val) {
                this.multipleSelection = val;
            },
            // 保存编辑
            saveEdit() {
                if (this.type == 0) {
                    updateUser({
                        Id: this.form.Id,
                        Username: this.form.Username,
                        Name: this.form.Name,
                        Password: this.form.Password,
                    }).then((res) => {
                        if (res.code == 0) {
                            this.editVisible = false;
                            this.$message.success(`修改第 ${this.idx + 1} 行成功`);
                            if (this.tableData[this.idx].Id === this.Id) {
                                this.$set(this.tableData, this.idx, this.form);
                            } else {
                                for (let i = 0; i < this.tableData.length; i++) {
                                    if (this.tableData[i].Id === this.Id) {
                                        this.$set(this.tableData, i, this.form);
                                        return;
                                    }
                                }
                            }
                        } else {
                            this.$message.error(`修改第 ${this.idx + 1} 行失败`);
                        }
                    });
                } else {
                    addUser({
                        Username: this.form.Username,
                        Name: this.form.Name,
                        Password: this.form.Password,
                    }).then((res) => {
                        if (res.code == 0) {
                            this.editVisible = false;
                            this.$message.success(`新增成功`);
                            this.getData();
                        } else {
                            this.$message.error(`新增失败`);
                        }
                    });
                }
            },
            // 确定删除
            deleteRow() {
                deleteUser({
                    ids: [this.id],
                }).then((res) => {
                    this.delVisible = false;
                    if (res.code == 0) {
                        this.$message.success('删除成功');
                        if (this.tableData[this.idx].Id === this.id) {
                            this.tableData.splice(this.idx, 1);
                        } else {
                            for (let i = 0; i < this.tableData.length; i++) {
                                if (this.tableData[i].Id === this.id) {
                                    this.tableData.splice(i, 1);
                                    return;
                                }
                            }
                        }
                    } else {
                        this.$message.error(res.message);
                    }
                });
            }
        }
    }

</script>

<style scoped>
    .handle-box {
        margin-bottom: 20px;
    }

    .handle-select {
        width: 120px;
    }

    .handle-input {
        width: 300px;
        display: inline-block;
    }

    .del-dialog-cnt {
        font-size: 16px;
        text-align: center
    }

    .table {
        width: 100%;
        font-size: 14px;
    }

    .red {
        color: #ff0000;
    }

    .mr10 {
        margin-right: 10px;
    }
</style>
