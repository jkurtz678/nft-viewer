import firebase from "../firebaseConfig";
import { FirestoreDocument, Account, BaseEntity } from "../types/types"
const db = firebase.firestore();

// returns a list of all accounts that match the given address (should only have 1 or 0 results unless something else has gone wrong)
export const getAccountByAddress = async (address: string): Promise<Array<FirestoreDocument<Account>>> => {
    const query_snapshot = await db.collection("account").where("address", "==", address).get();
    return query_snapshot.docs.map(s => ({
        id: s.id, entity: s.data() as Account,
    }))
};

// creates a new account record in the database and returns it
export const createAccount = async (address: string, signature: string): Promise<FirestoreDocument<Account>> => {
    const document = await db.collection("account")
        .add({
            ...BaseEntity.createBaseEntity(),
            address,
            signature
        })
    const snapshot = await document.get()
    return { id: snapshot.id, entity: snapshot.data() as Account }
}

export const updateAccount = async (account: FirestoreDocument<Account>): Promise<FirestoreDocument<Account>> => {
    await db.collection("account").doc(account.id).update(account.entity)
    const snapshot = await db.collection("account").doc(account.id).get()
    return { id: snapshot.id, entity: snapshot.data() as Account }
}
