/*
 * This file is part of Refractor.
 *
 * Refractor is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package rules

import (
	"Refractor/domain"
	"Refractor/params/validators"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"math"
)

var PlatformRules = []validation.Rule{
	validation.Length(1, 128),
	validation.By(validators.ValueInStrArray(domain.AllPlatforms)),
}

var PlayerIDRules = []validation.Rule{
	validation.Length(1, 80),
	is.Alphanumeric,
}

var InfractionReasonRules = []validation.Rule{
	validation.Length(1, 1024),
}

var InfractionDurationRules = []validation.Rule{
	validation.Min(0),
	validation.Max(math.MaxInt32),
}