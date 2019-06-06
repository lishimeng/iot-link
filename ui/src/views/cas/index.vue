<template>
  <div class="app-container">
    <el-row :gutter="30">
      <el-col :span="6" :xs="24">
        <cas-info :info="header" />
      </el-col>
    </el-row>
    <el-row :gutter="25">
      <el-col :span="12" :xs="24">
        <el-row>
          <el-col :span="24" :xs="24">
            <panel :title="base.title" :items="base.items" :shadow="always" />
          </el-col>
        </el-row>
      </el-col>
      <el-col :span="12" :xs="24">
        <el-row :gutter="15">
          <panel :title="character.title" :items="character.items" />
        </el-row>
        <el-row :gutter="15">
          <panel :title="security.title" :items="security.items" />
        </el-row>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import CasInfo from './components/CasInfo'
import Panel from './components/Panel'

export default {
  name: 'Cas',
  components: { CasInfo, Panel },
  data: () => {
    return {
      header: {
        label: '醋托啡',
        nameEN: 'Acetorphine',
        subLabel: '25333-77-1',
        formula: 'C19H21NO3',
        pic: 'https://gss0.bdstatic.com/94o3dSag_xI4khGkpoWK1HF6hhy/baike/s%3D220/sign=1993a5363fd3d539c53d08c10a86e927/d4628535e5dde7114096c63eaeefce1b9c1661b6.jpg'
      },
      base: {
        title: '基本信息',
        items: [
          { label: '中文名', value: '醋托啡' },
          { label: '英文名', value: 'Acetorphine' },
          { label: 'CAS号', value: '25333-77-1' },
          { label: '中文同义词', value: '烯丙吗啡' },
          { label: '英文同义词', value: '17-Allyl-7,8-didehydro-4,5alpha-epoxymorphinan-3,6alpha-diol;17-allyl-7,8-didehydro-4,5-alpha-epoxy-morphinan-6-alpha-diol;7,8-didehydro-4,5-epoxy-17-(2-propenyl)-,(5alpha,6alpha)-morphinan-6-diol;7,8-didehydro-4,5-epoxy-17-(2-propenyl)morphinan-3,6-diol;Acetorfin;Acetorphine;Allorphine;Anarcon' },
          { label: '备注', value: '-' }
        ]
      },
      security: {
        title: '安全信息',
        items: [
          { label: '药理特性', value: '-' },
          { label: '可燃性', value: '可燃' },
          { label: '危险类别', value: '中毒' },
          { label: '危险品运输编号', value: '-' }
        ]
      },
      character: {
        title: '物化特性',
        items: [
          { label: '分子式', value: 'C19H21NO3' },
          { label: '分子量g/mol', value: '311.37' },
          { label: '熔点℃', value: '208-209' },
          { label: '沸点℃', value: '451.42' }
        ]
      }
    }
  }
}
</script>

<style scoped>

</style>
