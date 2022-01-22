<template>
  <div>

    <h1 align="center" class="mb-12">ToDoの編集</h1>

    <v-row>
      <v-col cols="3"></v-col>
      <v-col cols="6">
        <v-text-field
          v-model="taskName"
          outlined
          label="タスクの変更"
          hide-selected
        ></v-text-field>
      </v-col>
      <v-col cols="3"></v-col>
    </v-row>

    <div align="center">
      <v-btn color="primary" @click="onChange()">変更</v-btn>
    </div>
    
  </div>
</template>

<script>
export default {
  name: "Edit",
  
  data() {
    return {
      taskName: null,
    }
  },
  methods: {
    async onChange() {
      let obj = {
        name: this.taskName
      }
      await this.axios({
        method: 'PUT',
        url: `/tasks/edit/${this.$route.params.id}`,
        data: obj
      }).then((res) => {
        console.log(res);
      }).catch((err) => {
        console.log(err);
      });

      this.$router.push(`/`);
    }
  },
  async created() {
    await this.axios({
      method: 'GET',
      url: `/tasks/${this.$route.params.id}`
    }).then((res) => {
      this.taskName = res.data[0].Name
      console.log(res);
    }).catch((err) => {
      console.log(err);
    });
  },
}
</script>