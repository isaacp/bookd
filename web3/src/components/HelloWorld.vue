<script>
// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
import { getFirestore } from "firebase/firestore";
import {collection, addDoc} from "firebase/firestore";

// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyB0AsbWcHMASMPyyVEv-HLq4aQLDLQ4vkE",
  authDomain: "isaacpendergrass-62b58.firebaseapp.com",
  projectId: "isaacpendergrass-62b58",
  storageBucket: "isaacpendergrass-62b58.appspot.com",
  messagingSenderId: "93576854239",
  appId: "1:93576854239:web:ca82c4894815cb390ac76a",
  measurementId: "G-MQW09JSB2T"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);
const db = getFirestore()


function write(who) {
}


import { reactive } from 'vue'
const gio = reactive({
  count: 0
})
const lucca = reactive({
  count: 0
})
const dom = reactive({
  count: 0
})
const otto = reactive({
  count: 0
})
const xander = reactive({
  count: 0
})
const max = reactive({
  count: 0
})
const oliver = reactive({
  count: 0
})
const sam = reactive({
  count: 0
})
const isaac = reactive({
  count: 0
})
const pender = reactive({
  count: 0
})
const deetrix = reactive({
  count: 0
})
const jared = reactive({
  count: 0
})
const ben = reactive({
  count: 0
})
const jayden = reactive({
  count: 0
})
</script>
<template>
  <v-container class="fill-height">
    <v-responsive class="align-center text-center fill-height">
      <!--<v-img height="300" src="@/assets/logo.svg" />-->

      <!--<div class="text-body-2 font-weight-light mb-n1">Welcome to</div>

      <h1 class="text-h2 font-weight-bold">Compliment Tracker</h1>-->

      <div class="py-14" />

      <v-row class="d-flex align-center justify-center">
        <v-col cols="auto">
          <v-btn
            color="primary"
            v-on:click = "gio.count--"
            min-width="35"
            target="_blank"
            variant="flat"
          >-</v-btn>
          <v-btn
            color="primary"
            v-on:click = "gio.count++"
            min-width="164"
            target="_blank"
            variant="flat"
          >
          <div style="display: flex; width: 100%; flex-direction: row; align-items: center; justify-content:space-evenly">
            <div>Gio</div>
          </div>
          </v-btn>
          <span style="background-color: white; color: black;padding: 8px; border-radius: 50px;">
              {{ gio.count }}
            </span>
        </v-col>

        <v-col cols="auto">
          <v-btn
            color="primary"
            v-on:click = "lucca.count--"
            min-width="35"
            target="_blank"
            variant="flat"
          >-</v-btn>
          <v-btn
            color="primary"
            v-on:click = "lucca.count++"
            min-width="164"
            target="_blank"
            variant="flat"
          >
          <div style="display: flex; width: 100%; flex-direction: row; align-items: center; justify-content:space-evenly">
            <div>Lucca</div>
          </div>
          </v-btn>
          <span style="background-color: white; color: black;padding: 8px; border-radius: 50px;">
              {{ lucca.count }}
            </span>
        </v-col>

        <v-col cols="auto">
          <v-btn
            color="primary"
            v-on:click = "dom.count--"
            min-width="35"
            target="_blank"
            variant="flat"
          >-</v-btn>
          <v-btn
            color="primary"
            v-on:click = "dom.count++"
            min-width="164"
            target="_blank"
            variant="flat"
          >
          <div style="display: flex; width: 100%; flex-direction: row; align-items: center; justify-content:space-evenly">
            <div>Dom</div>
          </div>
          </v-btn>
          <span style="background-color: white; color: black;padding: 8px; border-radius: 50px;">
              {{ dom.count }}
            </span>
        </v-col>

        <v-col cols="auto">
          <v-btn
            color="primary"
            v-on:click = "otto.count--"
            min-width="35"
            target="_blank"
            variant="flat"
          >-</v-btn>
          <v-btn
            color="primary"
            v-on:click = "otto.count++"
            min-width="164"
            target="_blank"
            variant="flat"
          >
          <div style="display: flex; width: 100%; flex-direction: row; align-items: center; justify-content:space-evenly">
            <div>Otto</div>
          </div>
          </v-btn>
          <span style="background-color: white; color: black;padding: 8px; border-radius: 50px;">
              {{ otto.count }}
            </span>
        </v-col>

        <v-col cols="auto">
          <v-btn
            color="primary"
            v-on:click = "sam.count--"
            min-width="35"
            target="_blank"
            variant="flat"
          >-</v-btn>
          <v-btn
            color="primary"
            v-on:click = "sam.count++"
            min-width="164"
            target="_blank"
            variant="flat"
          >
          <div style="displagioy: flex; width: 100%; flex-direction: row; align-items: center; justify-content:space-evenly">
            <div>sam</div>
          </div>
          </v-btn>
          <span style="background-color: white; color: black;padding: 8px; border-radius: 50px;">
              {{ sam.count }}
            </span>
        </v-col>

        <v-col cols="auto">
          <v-btn
            color="primary"
            v-on:click = "xander.count--"
            min-width="35"
            target="_blank"
            variant="flat"
          >-</v-btn>
          <v-btn
            color="primary"
            v-on:click = "xander.count++"
            min-width="164"
            target="_blank"
            variant="flat"
          >
          <div style="display: flex; width: 100%; flex-direction: row; align-items: center; justify-content:space-evenly">
            <div>xander</div>
          </div>
          </v-btn>
          <span style="background-color: white; color: black;padding: 8px; border-radius: 50px;">
              {{ xander.count }}
            </span>
        </v-col>

        <v-col cols="auto">
          <v-btn
            color="primary"
            v-on:click = "pender.count--"
            min-width="35"
            target="_blank"
            variant="flat"
          >-</v-btn>
          <v-btn
            color="primary"
            v-on:click = "pender.count++"
            min-width="164"
            target="_blank"
            variant="flat"
          >
          <div style="display: flex; width: 100%; flex-direction: row; align-items: center; justify-content:space-evenly">
            <div>pender</div>
          </div>
          </v-btn>
          <span style="background-color: white; color: black;padding: 8px; border-radius: 50px;">
              {{ pender.count }}
            </span>
        </v-col>

        <v-col cols="auto">
          <v-btn
            color="primary"
            v-on:click = "jared.count--"
            min-width="35"
            target="_blank"
            variant="flat"
          >-</v-btn>
          <v-btn
            color="primary"
            v-on:click = "jared.count++"
            min-width="164"
            target="_blank"
            variant="flat"
          >
          <div style="display: flex; width: 100%; flex-direction: row; align-items: center; justify-content:space-evenly">
            <div>jared</div>
          </div>
          </v-btn>
          <span style="background-color: white; color: black;padding: 8px; border-radius: 50px;">
              {{ jared.count }}
            </span>
        </v-col>

        <v-col cols="auto">
          <v-btn
            color="primary"
            v-on:click = "isaac.count--"
            min-width="35"
            target="_blank"
            variant="flat"
          >-</v-btn>
          <v-btn
            color="primary"
            v-on:click = "isaac.count++"
            min-width="164"
            target="_blank"
            variant="flat"
          >
          <div style="display: flex; width: 100%; flex-direction: row; align-items: center; justify-content:space-evenly">
            <div>isaac</div>
          </div>
          </v-btn>
          <span style="background-color: white; color: black;padding: 8px; border-radius: 50px;">
              {{ isaac.count }}
            </span>
        </v-col>

        <v-col cols="auto">
          <v-btn
            color="primary"
            v-on:click = "deetrix.count--"
            min-width="35"
            target="_blank"
            variant="flat"
          >-</v-btn>
          <v-btn
            color="primary"
            v-on:click = "deetrix.count++"
            min-width="164"
            target="_blank"
            variant="flat"
          >
          <div style="display: flex; width: 100%; flex-direction: row; align-items: center; justify-content:space-evenly">
            <div>deetrix</div>
          </div>
          </v-btn>
          <span style="background-color: white; color: black;padding: 8px; border-radius: 50px;">
              {{ deetrix.count }}
            </span>
        </v-col>

        <v-col cols="auto">
          <v-btn
            color="primary"
            v-on:click = "max.count--"
            min-width="35"
            target="_blank"
            variant="flat"
          >-</v-btn>
          <v-btn
            color="primary"
            v-on:click = "max.count++"
            min-width="164"
            target="_blank"
            variant="flat"
          >
          <div style="display: flex; width: 100%; flex-direction: row; align-items: center; justify-content:space-evenly">
            <div>max</div>
          </div>
          </v-btn>
          <span style="background-color: white; color: black;padding: 8px; border-radius: 50px;">
              {{ max.count }}
            </span>
        </v-col>

        <v-col cols="auto">
          <v-btn
            color="primary"
            v-on:click = "oliver.count--"
            min-width="35"
            target="_blank"
            variant="flat"
          >-</v-btn>
          <v-btn
            color="primary"
            v-on:click = "oliver.count++"
            min-width="164"
            target="_blank"
            variant="flat"
          >
          <div style="display: flex; width: 100%; flex-direction: row; align-items: center; justify-content:space-evenly">
            <div>oliver</div>
          </div>
          </v-btn>
          <span style="background-color: white; color: black;padding: 8px; border-radius: 50px;">
              {{ oliver.count }}
            </span>
        </v-col>

        <v-col cols="auto">
          <v-btn
            color="primary"
            v-on:click = "ben.count--"
            min-width="35"
            target="_blank"
            variant="flat"
          >-</v-btn>
          <v-btn
            color="primary"
            v-on:click = "ben.count++"
            min-width="164"
            target="_blank"
            variant="flat"
          >
          <div style="display: flex; width: 100%; flex-direction: row; align-items: center; justify-content:space-evenly">
            <div>ben</div>
          </div>
          </v-btn>
          <span style="background-color: white; color: black;padding: 8px; border-radius: 50px;">
              {{ ben.count }}
            </span>
        </v-col>

        <v-col cols="auto">
          <v-btn
            color="primary"
            v-on:click = "jayden.count--"
            min-width="35"
            target="_blank"
            variant="flat"
          >-</v-btn>
          <v-btn
            color="primary"
            v-on:click = "jayden.count++"
            min-width="164"
            target="_blank"
            variant="flat"
          >
          <div style="display: flex; width: 100%; flex-direction: row; align-items: center; justify-content:space-evenly">
            <div>jayden</div>
          </div>
          </v-btn>
          <span style="background-color: white; color: black;padding: 8px; border-radius: 50px;">
              {{ jayden.count }}
            </span>
        </v-col>
      </v-row>
    </v-responsive>
  </v-container>
</template>

<script setup>
  //
</script>
