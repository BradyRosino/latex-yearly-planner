# frozen_string_literal: true

module LatexYearlyPlanner
  module Planners
    module TypstMos
      module Pages
        class Face < Page
          def generate
            <<~TYPST
              #grid(
                columns: (#{heading_columns}),
                rows: (#{heading[:height]}, 1fr),
                grid.cell(
                  rowspan: 2,
                  pad(right: #{mosnav[:side_pad]}, #{side_menu_layout})
                ),
                pad(bottom: #{heading[:bottom_pad]}, #{header_layout}), #{content}
              )
            TYPST
          end

          def title
            raise NotImplementedError
          end

          def content
            raise NotImplementedError
          end

          def highlight_side_menu_quarters
            []
          end

          def side_menu_months
            []
          end

          def top_menu_month
            nil
          end

          def highlight_calendar?
            false
          end

          private

          def header_layout
            <<~TYPST
              table(
                columns: (auto, 1fr, auto),
                rows: 1fr,
                align: horizon + center,
                inset: 0mm,
                stroke: 0mm,
                #{title}, [], #{menu_items_layout}
              )
            TYPST
          end

          def menu_items_layout
            <<~TYPST
              table(
                  stroke: (x, y) => (left: 0.4pt, right: 0.4pt),
                  columns: #{menu_items_content.size},
                  rows: 1fr,
                  align: horizon + center,
                  #{menu_items_content.join(', ')}
                )
            TYPST
          end

          def menu_items_content
            @menu_items_content ||= [annual_menu_item].compact
          end

          def annual_menu_item
            return nil if !annual_section || !annual_section.enabled?

            name = i18n.t('menu_calendar')

            return "link(<annual-#{annual_page_number(top_menu_month)}>, [#{name}])" if top_menu_month
            return "table.cell(fill: black, link(<annual-1>, text(white)[#{name}]))" if highlight_calendar?

            "link(<annual-1>, [#{name}])"
          end

          def annual_section
            @annual_section ||= params.planner_config.sections.find { |section| section.name == :annual }
          end

          def annual_page_number(first_month)
            annual_params = params.planner_config.sections
                                  .find { |section| section.name == :annual }
                                  .params
            months_per_page = annual_params.get(:months_per_page)

            (annual_params.months.find_index(first_month) / months_per_page) + 1
          end

          def side_menu_layout
            <<~TYPST
              rotate(
                #{mosnav[:rotate]},
                origin: center + horizon,
                reflow: true,
                table(
                  stroke: (x, y) => (left: 0.4pt, right: 0.4pt),
                  columns: (#{side_menu_columns}),
                  rows: 1fr,
                  align: horizon + center,
                  #{side_menu_content}
                )
              )
            TYPST
          end

          def side_menu_columns
            cols = ([mosnav[:quarter_width]] * params.quarters.size)
                   .append('auto')
                   .append([mosnav[:month_width]] * params.months.size)

            cols.reverse! if mosnav[:reverse_arrays]

            cols.join(', ')
          end

          def side_menu_content
            quarters = make_side_menu_quarters
            months = make_side_menu_months

            if mosnav[:reverse_array_internals]
              quarters.reverse!
              months.reverse!
            end

            side_menu = [quarters.join(', '), '[]', months.join(', ')]

            side_menu.reverse! if mosnav[:reverse_arrays]

            side_menu.join(', ')
          end

          def make_side_menu_quarters
            params.quarters.map do |q|
              name = i18n.t('calendar.one_letter.quarter')

              next "link(<#{q.id}>, [#{name}#{q.number}])" unless highlight_side_menu_quarters.include?(q)

              "table.cell(fill: black, link(<#{q.id}>, text(white)[#{name}#{q.number}]))"
            end
          end

          def make_side_menu_months
            params.months.map do |m|
              name = i18n.t("calendar.short.month.#{m.name.downcase}")

              next "link(<#{m.id}>, [#{name}])" unless side_menu_months.include?(m)

              "table.cell(fill: black, link(<#{m.id}>, text(white)[#{name}]))"
            end
          end

          def heading_columns
            "#{mosnav[:width]}, 1fr"
          end

          def mosnav
            params.object(:mos_nav)
          end

          def heading
            params.object(:heading)
          end
        end
      end
    end
  end
end
