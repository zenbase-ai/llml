from ..base.bool import format_bool, is_bool
from ..base.datetime import format_datetime, is_datetime
from ..base.fallback import format_fallback, is_fallback
from ..base.float import format_float, is_float
from ..base.int import format_int, is_int
from ..base.none import format_none, is_none
from ..base.str import format_str, is_str
from ..base.timedelta import format_timedelta, is_timedelta
from ..types import Formatters
from .dict import format_dict, is_dict
from .list import format_list, is_list

swag_xml: Formatters = {
    is_str: format_str,
    is_int: format_int,
    is_float: format_float,
    is_bool: format_bool,
    is_none: format_none,
    is_datetime: format_datetime,
    is_timedelta: format_timedelta,
    is_list: format_list,
    is_dict: format_dict,
    is_fallback: format_fallback,
}
