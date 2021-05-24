// src/firebaseConfig.js
import firebase from "firebase/app";

// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyAdu0sAk8jWlbjGyqHpULmu4E8modXob_8",
  authDomain: "nft-viewer.firebaseapp.com",
  projectId: "nft-viewer",
  storageBucket: "nft-viewer.appspot.com",
  messagingSenderId: "680051604339",
  appId: "1:680051604339:web:aa0793be255390e752e306",
  measurementId: "G-8PLTH0Z68H"
};

// Initialize Firebase
export default firebase.initializeApp(firebaseConfig);