<template>
  <div>
    <a-form :form="form" @submit="handleSubmit">

      <a-form-item
        v-show="isEdit"
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        label="任务ID"
        hasFeedback
      >
        <a-input
          placeholder="任务ID"
          v-decorator="[
            'id',
            {rules: []}
          ]"
          :disabled="true"
        ></a-input>
      </a-form-item>
      <a-form-item
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        label="备份周期"
        hasFeedback
        validateStatus="success"
      >
        <a-input style="width: 100%" placeholder="请输入备份周期 ex：30 12 * * *" v-decorator="['backup_cycle', {rules: [{ required: true }]}]" />
      </a-form-item>

      <a-form-item
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        label="数据保留周期"
        hasFeedback
        validateStatus="success"
      >
        <a-input-number :min="1" style="width: 100%" placeholder="请输入数据保留周期(天)" v-decorator="['keep_number', {rules: [{ required: true }]}]" />
      </a-form-item>
      <a-form-item
        v-bind="buttonCol"
      >
        <a-row>
          <a-col span="6">
            <a-button type="primary" html-type="submit">提交</a-button>
          </a-col>
          <a-col span="10">
            <a-button @click="handleGoBack">返回</a-button>
          </a-col>
          <a-col span="8"></a-col>
        </a-row>
      </a-form-item>
    </a-form>
  </div>
</template>

<script>

import pick from 'lodash.pick'
import { addEsTask, updateEsTask } from '@/api/elastic'
import { GetAgentTaskDetail } from '@/api/agent-task'

export default {
  name: 'TableEdit',
  props: {
    record: {
      type: [Object, String],
      default: ''
    },
    host: {
      type: Number,
      default: 0
    }
  },
  data () {
    return {
      labelCol: {
        xs: { span: 24 },
        sm: { span: 5 }
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 12 }
      },
      buttonCol: {
        wrapperCol: {
          xs: { span: 24 },
          sm: { span: 12, offset: 5 }
        }
      },
      form: this.$form.createForm(this),
      id: 0,
      isEdit: false,
      // 开关状态相关
      dingStatus: false,
      ossStatus: false,
      // 存储类型相关
      ossType: 1,
      // host_id,用于编辑时传递
      hostID: 0
    }
  },
  // beforeCreate () {
  //   this.form = this.$form.createForm(this)
  // },
  created () {
    this.service_name = this.$route.params && this.$route.params.service_name
  },
  mounted () {
    this.$nextTick(() => {
      this.loadEditInfo(this.record)
    })
  },
  methods: {
    handleGoBack () {
      this.$emit('onGoBack')
    },
    handleSubmit (e) {
      e.preventDefault()
      // const { form: { validateFields } } = this
      this.form.validateFields((err, values) => {
        console.log(err)
           if (!err) {
             if (this.isEdit) {
               values['host_id'] = this.host
               values['service_name'] = this.service_name
               updateEsTask(values).then((res) => {
                 if (res) {
                   this.$message.success(res.data)
                   this.handleGoBack()
                 } else {
                   this.$message.error('修改失败')
                 }
               })
             } else {
               values['host_id'] = this.host
               values['service_name'] = this.service_name
               addEsTask(values).then((res) => {
                 if (res) {
                   this.$message.success(res.data)
                   this.handleGoBack()
                 }
               })
             }
           }
      })
    },
    BoolToInt (i) {
      if (i) {
        return 1
      } else {
        return 0
      }
    },
    handleGetInfo () {

    },
    loadEditInfo (data) {
      const { form } = this
      if (data.id) {
        this.isEdit = true
        new Promise((resolve) => {
          resolve()
        }).then(() => {
          const query = { 'id': data.id, 'service_name': this.service_name }
          GetAgentTaskDetail(query).then((res) => {
            const formData = pick(res.data, ['host', 'host_id', 'backup_cycle', 'db_name', 'status', 'keep_number', 'id',
              'is_ding_send', 'ding_access_token', 'ding_secret',
            'is_oss_save', 'oss_type', 'endpoint', 'oss_access', 'oss_secret', 'bucket_name', 'directory'
            ])
            this.dingStatus = Boolean(formData.is_ding_send)
            this.ossStatus = Boolean(formData.is_oss_save)
            this.hostID = formData.host_id
            // 如果开关打开,修改存储类型
            if (formData.is_oss_save !== 0) {
              this.ossType = formData.oss_type
            }
            form.setFieldsValue(formData)
          })
        })
      }
    }
  }
}
</script>
