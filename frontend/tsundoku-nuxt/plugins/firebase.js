import firebase from 'firebase/app';
import 'firebase/auth';

if (!firebase.apps.length) {
    firebase.initializeApp({
        apiKey: 'AIzaSyC_oob5ylts7a7-Zfcq0xEtdYBs81bLHhk',
        authDomain: 'tsundoku-5c328.firebaseapp.com',
        projectId: 'tsundoku-5c328',
        storageBucket: 'tsundoku-5c328.appspot.com',
        messagingSenderId: '607045266249',
        appId: '1:607045266249:web:df18f616ea3e51687b0898',
        measurementId: 'G-R26SH865FQ'
    });
}

export default firebase;
