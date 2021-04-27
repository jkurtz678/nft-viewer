import firebase from "firebase";

export interface IBaseEntity {
    created_at: firebase.firestore.Timestamp;
    updated_at: firebase.firestore.Timestamp;
}

export class BaseEntity {

    static createBaseEntity(): IBaseEntity {
        const now = firebase.firestore.Timestamp.now();

        return { created_at: now, updated_at: now };
    }

    static updateBaseEntity(): Partial<IBaseEntity> {
        const now = firebase.firestore.Timestamp.now();

        return { updated_at: now }
    }
}

export interface EntityPayload<Entity extends IBaseEntity> {
    entity: Entity;
    id: string;
}

export interface Account extends IBaseEntity {
    address: string;
    signature: string;
    device_ids: Array<string>;
}