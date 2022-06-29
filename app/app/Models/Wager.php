<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;
use Illuminate\Support\Carbon;

/**
 * App\Models\Wager
 *
 * @property int $id
 * @property float $total_wager_value
 * @property float $odds
 * @property float $selling_percentage
 * @property float $selling_price
 * @property float $current_selling_price
 * @property float $percentage_sold
 * @property float $amount_sold
 * @property string $placed_at
 * @property Carbon|null $created_at
 * @property Carbon|null $updated_at
 * @property Carbon|null $deleted_at
 * @method static \Illuminate\Database\Eloquent\Builder|Wager newModelQuery()
 * @method static \Illuminate\Database\Eloquent\Builder|Wager newQuery()
 * @method static \Illuminate\Database\Query\Builder|Wager onlyTrashed()
 * @method static \Illuminate\Database\Eloquent\Builder|Wager query()
 * @method static \Illuminate\Database\Eloquent\Builder|Wager whereAmountSold($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Wager whereCreatedAt($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Wager whereCurrentSellingPrice($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Wager whereDeletedAt($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Wager whereId($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Wager whereOdds($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Wager wherePercentageSold($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Wager wherePlacedAt($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Wager whereSellingPercentage($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Wager whereSellingPrice($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Wager whereTotalWagerValue($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Wager whereUpdatedAt($value)
 * @method static \Illuminate\Database\Query\Builder|Wager withTrashed()
 * @method static \Illuminate\Database\Query\Builder|Wager withoutTrashed()
 * @mixin \Eloquent
 */
class Wager extends Model
{
    use SoftDeletes;

    protected $table = 'wagers';

    protected $fillable = [
        'total_wager_value',
        'odds',
        'selling_percentage',
        'selling_price',
        'current_selling_price',
        'amount_sold',
        'placed_at',
    ];
}
