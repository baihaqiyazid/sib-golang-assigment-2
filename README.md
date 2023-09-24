# sib-golang-assigment-2

# Result

## Create Order
![image](https://github.com/baihaqiyazid/sib-golang-assigment-2/assets/89854394/a72d2de4-83a1-45be-9812-f45f5dc71b1c)

```
{
    "code": 200,
    "status": "success",
    "data": {
        "id": 8,
        "customer_name": "Baihaqi",
        "items": [
            {
                "item_id": 45,
                "item_code": "ikl-mod12",
                "description": "Hoodie H&M",
                "quantity": 2,
                "order_id": 8,
                "created_at": "2023-09-25T06:03:27.533168358+07:00",
                "updated_at": "2023-09-25T06:03:27.533168358+07:00"
            },
            {
                "item_id": 46,
                "item_code": "ikl-mod4",
                "description": "Iphone 11",
                "quantity": 1,
                "order_id": 8,
                "created_at": "2023-09-25T06:03:27.533168358+07:00",
                "updated_at": "2023-09-25T06:03:27.533168358+07:00"
            }
        ],
        "ordered_at": "2019-11-09T21:21:46Z",
        "created_at": "2023-09-25T06:03:27.431012029+07:00",
        "updated_at": "2023-09-25T06:03:27.431012207+07:00"
    }
}
```

## Get All Order
![image](https://github.com/baihaqiyazid/sib-golang-assigment-2/assets/89854394/179911e3-6208-41f1-814c-07777f1fbe3c)

```
{
    "code": 200,
    "status": "success",
    "data": [
        {
            "id": 1,
            "customer_name": "Jakcir",
            "items": [
                {
                    "item_id": 1,
                    "item_code": "12312323asdw",
                    "description": "Iphone 10x",
                    "quantity": 1,
                    "order_id": 1,
                    "created_at": "2023-09-23T13:59:28.188704+07:00",
                    "updated_at": "2023-09-23T13:59:28.188704+07:00"
                },
                {
                    "item_id": 2,
                    "item_code": "12312323asdwsad",
                    "description": "Iphone 14",
                    "quantity": 1,
                    "order_id": 1,
                    "created_at": "2023-09-23T13:59:28.188704+07:00",
                    "updated_at": "2023-09-23T13:59:28.188704+07:00"
                }
            ],
            "ordered_at": "2019-11-10T04:21:46+07:00",
            "created_at": "2023-09-23T13:59:28.086465+07:00",
            "updated_at": "2023-09-23T16:41:13.073538+07:00"
        },
        {
            "id": 7,
            "customer_name": "Liam Nesson",
            "items": [
                {
                    "item_id": 43,
                    "item_code": "12312323asdw",
                    "description": "Iphone 10x",
                    "quantity": 3,
                    "order_id": 7,
                    "created_at": "2023-09-24T09:53:30.890265+07:00",
                    "updated_at": "2023-09-24T09:53:30.890265+07:00"
                },
                {
                    "item_id": 44,
                    "item_code": "12312323asdwsad",
                    "description": "Iphone 14",
                    "quantity": 5,
                    "order_id": 7,
                    "created_at": "2023-09-24T09:53:30.890265+07:00",
                    "updated_at": "2023-09-24T09:53:30.890265+07:00"
                }
            ],
            "ordered_at": "2019-11-10T04:21:46+07:00",
            "created_at": "2023-09-24T09:53:30.788832+07:00",
            "updated_at": "2023-09-24T09:53:30.788832+07:00"
        }
    ]
}
```

## Update Order
![image](https://github.com/baihaqiyazid/sib-golang-assigment-2/assets/89854394/bb8446a5-b630-420e-b8f6-ec6c2f06739a)

```
{
    "code": 200,
    "status": "success",
    "data": {
        "id": 8,
        "customer_name": "Baihaqi Yazid",
        "items": [
            {
                "item_id": 45,
                "item_code": "ikl-mod12",
                "description": "Samsung A50",
                "quantity": 2,
                "order_id": 8,
                "created_at": "2023-09-25T06:03:27.533168+07:00",
                "updated_at": "2023-09-25T06:06:28.613345+07:00"
            },
            {
                "item_id": 46,
                "item_code": "ikl-mod15",
                "description": "T-shirt Polo",
                "quantity": 4,
                "order_id": 8,
                "created_at": "2023-09-25T06:03:27.533168+07:00",
                "updated_at": "2023-09-25T06:06:28.618183+07:00"
            }
        ],
        "ordered_at": "2019-11-10T04:21:46+07:00",
        "created_at": "2023-09-25T06:03:27.431012+07:00",
        "updated_at": "2023-09-25T06:06:28.6216+07:00"
    }
}
```

## Delete Order

![image](https://github.com/baihaqiyazid/sib-golang-assigment-2/assets/89854394/a939504f-0ab1-46e6-9f1e-c6354f821834)



