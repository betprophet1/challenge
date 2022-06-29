<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;
use Illuminate\Support\Carbon;

/**
 * App\Models\Purchase
 *
 * @property int $id
 * @property int $wager_id
 * @property float $buying_price
 * @property string $bought_at
 * @property Carbon|null $created_at
 * @property Carbon|null $updated_at
 * @property Carbon|null $deleted_at
 * @method static \Illuminate\Database\Eloquent\Builder|Purchase newModelQuery()
 * @method static \Illuminate\Database\Eloquent\Builder|Purchase newQuery()
 * @method static \Illuminate\Database\Query\Builder|Purchase onlyTrashed()
 * @method static \Illuminate\Database\Eloquent\Builder|Purchase query()
 * @method static \Illuminate\Database\Eloquent\Builder|Purchase whereBoughtAt($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Purchase whereBuyingPrice($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Purchase whereCreatedAt($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Purchase whereDeletedAt($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Purchase whereId($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Purchase whereUpdatedAt($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Purchase whereWagerId($value)
 * @method static \Illuminate\Database\Query\Builder|Purchase withTrashed()
 * @method static \Illuminate\Database\Query\Builder|Purchase withoutTrashed()
 * @mixin \Eloquent
 */
class Purchase extends Model
{
    use SoftDeletes;

    protected $table = 'purchases';

    protected $fillable = [
        'wager_id',
        'buying_price',
        'bought_at',
    ];
}
