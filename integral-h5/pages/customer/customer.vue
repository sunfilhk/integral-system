<template>
  <view class="container">
    <view class="searchBox">
      <uni-easyinput
        trim="all"
        @input="searchInput"
        v-model="queryParams.keyword"
        placeholder="请输入用户名/手机号搜索"
        suffixIcon="search"
      />
    </view>
    <button class="newcustomers" size="mini" type="primary" @click="customer">
      新增客户
    </button>
    <view
      class="tab"
      v-for="(item, index) in list"
      :key="index"
      @click="infor(item)"
    >
      <view>客户姓名：{{ item.name }}</view>
      <view>创建时间：{{ parseTime(item.create_time) }}</view>
      <view class="status"
        >状态：
        <text class="a" v-show="item.status == 0">已通过</text>
        <text class="b" v-show="item.status == 1">审核中</text>
        <text class="c" v-show="item.status == 3">已拒绝</text>
      </view>
      <view class="btn">
        <view
          ><button
            size="mini"
            type="default"
            @click.stop="investment(item.id, item.flag)"
          >
            投资列表
          </button></view
        >
      </view>
    </view>
    <view v-show="isLoadMore">
      <uni-load-more :status="loadStatus"></uni-load-more>
    </view>
  </view>
</template>

<script>
import { customerList } from '../../common/api.js'
export default {
  data () {
    return {
      list: [],
      queryParams: {
        keyword: '', //string类型 用户名或手机号模糊查询
        pageIndex: 1,
        pageSize: 10
      },
      loadStatus: 'loading', //加载样式：more-加载前样式，loading-加载中样式，nomore-没有数据样式
      isLoadMore: false //是否加载中
    }
  },
  onShow () {
    this.list = []
    this.auto()
  },
  onReachBottom () {
    //上拉触底函数
    if (!this.isLoadMore) {
      //此处判断，上锁，防止重复请求
      this.isLoadMore = true
      this.queryParams.pageIndex += 1
      this.auto()
    }
  },
  onPullDownRefresh () {
    this.list = []
    this.queryParams.pageIndex = 1
    this.auto()
    uni.stopPullDownRefresh()
  },
  methods: {
    customer () {
      uni.navigateTo({
        url: '/pages/incustomer/incustomer'
      })
    },
    auto (search) {
      customerList(this.queryParams).then(res => {
        if (res.list) {
          if (search) {
            this.list = res.list
            if (res.list.length < this.queryParams.pageSize) {
              //判断接口返回数据量小于请求数据量，则表示此为最后一页
              this.isLoadMore = true
              this.loadStatus = 'nomore'
            } else {
              this.isLoadMore = false
            }
          } else {
            this.list = this.list.concat(res.list)
            if (res.list.length < this.queryParams.pageSize) {
              //判断接口返回数据量小于请求数据量，则表示此为最后一页
              this.isLoadMore = true
              this.loadStatus = 'nomore'
            } else {
              this.isLoadMore = false
            }
          }
        } else {
          this.isLoadMore = true
          this.loadStatus = 'nomore'
        }
      })
    },
    infor (item) {
      uni.navigateTo({
        url: '/pages/infor/infor?id=' + item.id + '&flag=' + item.flag
      })
    },
    investment (id, flag) {
      uni.navigateTo({
        url: '/pages/investment/investment?id=' + id + '&flag=' + flag
      })
    },
    // 搜索方法
    searchInput (e) {
      this.queryParams = {
        keyword: e,
        pageIndex: 1,
        pageSize: 10
      }

      this.auto('search')
    }
  }
}
</script>

<style>
.newcustomers {
  margin: 20px 20px 10px;
}

.tab {
  font-size: 14px;
  color: #444444;
  border-top: 1px solid #dddddd;
  padding: 11px 20px 15px;
  position: relative;
}

.tab view {
  margin-top: 4px;
}

.tab .status .a {
  color: #00bc00;
}

.tab .status .b {
  color: #007aff;
}

.tab .status .c {
  color: rgb(255, 80, 80);
}

.btn {
  position: absolute;
  right: 10px;
  top: 24px;
}

.btn button::after {
  border: none;
}

.btn button {
  background: #ffffff;
  color: #007aff;
}
/* 搜索框 */
.searchBox {
  padding: 10px 20px 0px;
}
</style>
