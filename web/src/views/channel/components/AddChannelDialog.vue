<template>
	<el-dialog
		:title="dataSource ? '编辑' : '创建'"
		:visible.sync="dialogVisible"
		width="620px"
		custom-class="wk-dialog"
		append-to-body
		center
		:close-on-click-modal="false"
		@open="handleOpen"
		:before-close="handleClose"
	>
		<div class="form">
			<el-form class="wk-form" ref="form" :model="form" :rules="rules" label-width="140px">
				<el-form-item label="通道名称：" prop="name">
					<el-input class="wk-input" size="small" v-model="form.name"></el-input>
				</el-form-item>
        <el-form-item label="通道流地址：" prop="url">
					<el-input class="wk-input" size="small" v-model="form.url"></el-input>
				</el-form-item>
        <el-form-item label="通道来源：" prop="source">
					<el-select v-model="form.source" placeholder="请选择">
						<el-option v-for="item in Object.keys(CHANNEL_TYPE)" :key="item" :label="CHANNEL_TYPE[item].label" :value="CHANNEL_TYPE[item].value"></el-option>
					</el-select>
					<!-- <el-input class="wk-input" size="small" v-model="form.source"></el-input> -->
				</el-form-item>
        <el-form-item label="通道编码：" prop="code">
					<el-input class="wk-input" size="small" v-model="form.code"></el-input>
				</el-form-item>
        <el-form-item label="是否转封装：" prop="repackage">
					<el-select v-model="form.repackage" placeholder="请选择">
						<el-option v-for="item in Object.keys(REPACKAGE_TYPE)" :key="item" :label="REPACKAGE_TYPE[item].label" :value="REPACKAGE_TYPE[item].value"></el-option>
					</el-select>
					<!-- <el-input class="wk-input" size="small" v-model="form.repackage"></el-input> -->
				</el-form-item>
        <el-form-item label="封装协议：" prop="repackageFormat">
					<el-input class="wk-input" size="small" v-model="form.repackageFormat"></el-input>
				</el-form-item>
			</el-form>
		</div>
		<div slot="footer" class="dialog-footer">
			<el-button size="small" class="wk-default-button" @click="handleClose">取 消</el-button>
			<el-button size="small" class="wk-primary-button" type="primary" @click="submit">确 定</el-button>
		</div>
	</el-dialog>
</template>
<script>
import { cloneDeep } from 'lodash';
import { CHANNEL_TYPE, REPACKAGE_TYPE } from "./../config";

export default {
	props: ['dialogVisible', 'handleClose', 'dataSource'],
	data() {
		return {
			CHANNEL_TYPE,
			REPACKAGE_TYPE,
			form: {
				name: '',
				url: '',
        source: '',
        code: '',
        gbCode: '',
        repackage: '',
        repackageFormat: ''
			},
			cacheForm: null,
			rules: {
				name: [{ required: true, message: '请输入模库名称', trigger: 'blur' }],
			},
			isEdit: false,
		};
	},
	mounted() {
		this.cacheForm = cloneDeep(this.form);
	},
	methods: {
		handleOpen() {
			this.$nextTick(() => {
				this.$refs.form.resetFields();
			});
			if (this.dataSource) {
				this.isEdit = true;
				this.form = cloneDeep(this.dataSource);
			} else {
				this.isEdit = false;
				this.form = cloneDeep(this.cacheForm);
			}
		},
		async submit() {
			const valid = await this.$refs.form.validate();
			if (!valid) return false;
			const params = cloneDeep(this.form);
			if (this.isEdit) {
			}
		},
	},
};
</script>
<style lang='scss' scoped>
</style>