<template>
  <div>
    <el-dialog :title="title"
               v-model="visible" size="tiny">
      <img v-if="imgUrl"
           :src="imgUrl"
           class="pic-preview">
    </el-dialog>
    <!--<input type="text"
               :value="value"
               style='display:none;'>-->
  </div>
</template>
<script>
export default {
  props: {
    imgUrl: {
      type: null,
      required: true
    },
    value: {
      type: Boolean,
      default: false
    },
    title: {
      type: String,
    }
  },
  data: function () {
    return {
      visible: false
    }
  },
  watch: {
    value(val) {
      if (this.imgUrl === '') {
        this.$message.warning('没有图片')
        this.$emit('input', false);
        return
      }
      this.visible = val;
    },
    visible(val) {
      this.$emit('input', val);
    }
  },
  mounted() {
    if (this.value) {
      this.visible = true;
    }
  },
//   computed: {
//     imgUrl: function () {
//       return config.env.API_ROOT + api.version + api.asset + '/' + this.picId + '?_token=' + this.$store.state.token;
//     }
//   },
}
</script>
<style>
.pic-preview {
  width: 100%;
  height: 100%;
  display: block;
}
</style>