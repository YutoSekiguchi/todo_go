<template>
  <div>

    <h1 align="center">ToDo</h1>

    <v-row>
      <v-col cols="4"></v-col>
      <v-col cols="4">
        <v-text-field
          v-model="taskName"
          outlined
          dense
          label="タスク名"
          hide-selected
        ></v-text-field>
      </v-col>
      <v-col cols="4" align="start">
        <v-btn @click="onSubmit()">登録</v-btn>
      </v-col>
    </v-row>

    <v-row 
      v-for="n in taskList.length" 
      :key="n"
    >
      <v-col cols="2"></v-col>
      <v-col align="start" cols="8" class="todolist">
        <v-row>
          <v-col><h2>{{ taskList[n-1].Name }}</h2></v-col>
          <v-col align="end">
            <v-btn color="success" class="mr-4" @click="handleMoveToEdit(taskList[n-1].ID)">編集</v-btn>
            <v-btn color="error" @click="onDelete(taskList[n-1].ID)">削除</v-btn>
          </v-col>
        </v-row>
      </v-col>
      <v-col cols="2"></v-col>
    </v-row>

  </div>
</template>

<script>


  export default {
    name: 'Home',

    data() {
      return {
        taskList: [], // タスク一覧を入れるリスト
        taskName: null, // タスクの登録時に使用するタスク名
      }
    },
    methods: {
      async onSubmit() { // ボタンを押した時の処理
        if (this.taskName != null) {
          let obj = {
            name: this.taskName
          }
          await this.axios({ // taskの追加
            method: 'POST',
            url: `/tasks`,
            data: obj
          }).then((res) => {
            console.log(res);
          }).catch((error) => {
            console.log(error);
          });

          this.taskName = null; // formの中身を空にするためtaskNameをnullに
          this.readAllTask();
        }
      },
      // 編集ボタンを押した時の処理
      async handleMoveToEdit(id) {
        await this.$router.push(`/edit/${id}`);
      },
      // 削除ボタンを押した時の処理
      async onDelete(id) {
        await this.axios({
          method: 'DELETE',
          url: `/tasks/delete/${id}`
        }).then((res) => {
          console.log(res);
        }).catch((error) => {
          console.log(error);
        });
        this.readAllTask();
      },
      async readAllTask() { // 全てのタスクを呼び出してtaskListに追加
        await this.axios({ 
          method: 'GET',
          url: `/tasks`,
        }).then((res) => {
          this.taskList = res.data;
        }).catch((error) => {
          console.log(error);
        });
      }
    },
    created() {
      this.readAllTask()
    }
  }
</script>

<style scoped>
.todolist {
  border: 1px solid grey;
  border-radius: 10px;
  margin-bottom: 20px;
}
</style>