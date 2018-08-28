<template>
<div class="exsample">
  <img v-bind:src="members[counter].Img"/>
  <h1 v-on:click="handleTitleOpen">{{ isTitleOpen ? members[counter].Title : "役職"}}</h1>
  <h1 v-on:click="handleNameOpen">{{ isNameOpen ? members[counter].Name : "ニックネーム"}}</h1>
  <button v-on:click="nextMember">次へ</button>

</div>
</template>

<script lang="ts">
    import { Component, Prop, Vue } from 'vue-property-decorator';
import axios from 'axios';


interface Member {
    Img:string
    Title:string
    Name:string
}

@Component
export default class Exsmple extends Vue {
    @Prop() private str!: string;

    private members: Member[] = [{Img:"hogehoge",Title:"",Name:"ロード中"}]
    private counter: number = 0;
    private isTitleOpen: boolean = false;
    private isNameOpen: boolean = false;

    private mounted(): void {
	axios
	    .get('api/member/get')
	    .then(res => {
		this.members = res.data;
	    })
	    .catch(err => {
		console.log(err)
	    });
    }

    private handleTitleOpen(): void {
	this.isTitleOpen = true;
    }

    private handleNameOpen(): void {
	this.isNameOpen = true;
    }

    private nextMember(): void {
	if (this.counter !== this.members.length - 1){
	    this.counter++;
	}else{
	    this.counter = 0;
	};
	this.isTitleOpen = false;
	this.isNameOpen = false;
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
  .exsample {
  background-color: #888;
  }

  h1 {
  cursor:pointer;
  }

</style>
