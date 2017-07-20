<?php
/**
 * User: WindomZ
 * Date: 17-7-20
 */

namespace GraphQLExamples\Field;

use GraphQL\Type\Definition\ObjectType;
use GraphQL\Type\Definition\ResolveInfo;
use GraphQL\Type\Definition\Type;

class User
{
    public $id;

    public $name;

    public function __construct($id, $name)
    {
        $this->id = $id;
        $this->name = $name;
    }

    private static $data;

    private static function getData()
    {
        return self::$data ?: (self::$data = [
            '1' => new User('1', 'Name-1'),
            'id' => new User('id', 'Name'),
            '编号' => new User('编号', '名字'),
        ]);
    }


    public static function getUserField()
    {
        return [
            'name' => 'user',
            'type' => UserType::UserType(),
            'args' => [
                'id' => ['type' => Type::string()],
            ],
            'resolve' => function ($root, $args) {
                return self::getData()[$args['id']];
            },
        ];
    }
}

class UserType extends ObjectType
{
    private static $user;

    /**
     * @return UserType
     */
    public static function UserType()
    {
        return self::$user ?: (self::$user = new UserType());
    }

    public function __construct()
    {
        $config = [
            'name' => 'User',
            'description' => 'Our blog authors',
            'fields' => function () {
                return [
                    'id' => Type::string(),
                    'name' => Type::string(),
                ];
            },
            'resolveField' => function ($value, $args, $context, ResolveInfo $info) {
                return $value->{$info->fieldName};
            }
        ];
        parent::__construct($config);
    }
}

