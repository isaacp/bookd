<script setup>
import TheWelcome from '../components/TheWelcome.vue'
import axios from 'axios'
// Import the functions you need from the SDKs you need
import { initializeApp } from "https://www.gstatic.com/firebasejs/10.8.0/firebase-app.js";
import { getAnalytics } from "https://www.gstatic.com/firebasejs/10.8.0/firebase-analytics.js";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyDnhRwd24usMMrqKtD9I5oH_wp_F20e6bM",
  authDomain: "bookd-656ca.firebaseapp.com",
  projectId: "bookd-656ca",
  storageBucket: "bookd-656ca.appspot.com",
  messagingSenderId: "415801325888",
  appId: "1:415801325888:web:f7f15bf79b336a969fa2a5",
  measurementId: "G-JSVTBQC1H6"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);
const start = defineModel('start')
const end = defineModel('end')
</script>

<template>
  <input v-model="start"/>  <input v-model="end"/> <button @click="getData(start, end)"></button>
  <br/>
  <br/>
  <div v-for="event in events" style="padding-bottom: 25px">
   {{ formatDate(event.start, "date") + " ( " + formatDate(event.start, "time") + " - " + formatDate(event.end, "time") + " )"}}
   <br/>
   {{ event.summary }}
  <br/>
   {{ event.location }}
  </div>
</template>
<script>
export default {
  data() {
    return {
      events: []
    }
  },
  mounted() {
    let now = new Date()
    let then = new Date()
    then.setDate(then.getDate() + 60)
    axios.get("https://api-bikzyn25da-uw.a.run.app/api/events?start="+ now.toISOString() +"&end=" + then.toISOString()).then(response => {
      this.events = response.data
    })
  },
  methods: {
    getData(start, end) {
      axios.get("https://api-bikzyn25da-uw.a.run.app/api/events?start="+start+"&end="+end).then(response => {
      this.events = response.data
    })
    },
    formatDate(value, time = "") {
      let options = {
          year: "numeric",
          month: "numeric",
          day: "numeric",
          hour: "numeric",
          minute: "numeric",
          hour12: false,
          timeZone: "America/Los_Angeles",
        };
        
      if (time == "time") {
        options = {
          hour: "numeric",
          minute: "numeric",          
          hour12: false,
          timeZone: "America/Los_Angeles",
        };
      } else if (time == "date") {
        options = {
          year: "numeric",
          month: "numeric",
          day: "numeric",
          hour12: false,
          timeZone: "America/Los_Angeles",
        };
      }
      const date = new Date(value);
      // Then specify how you want your dates to be formatted
      return new Intl.DateTimeFormat('default', options).format(date);
    }
  }
};
</script>
